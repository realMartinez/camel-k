Feature: Camel K can load default secrets for Kamelets

  Background:
    Given Disable auto removal of Kamelet resources
    Given Disable auto removal of Kubernetes resources
    Given Camel K resource polling configuration
      | maxAttempts          | 40   |
      | delayBetweenAttempts | 3000 |

  Scenario: Pipe can load default settings for Kamelet
    Given create Kubernetes service stub-service with target port 8080
    And bind Kamelet timer-source to uri http://stub-service.${YAKS_NAMESPACE}.svc.cluster.local/default
    When create Pipe binding
    Then Pipe binding should be available

 Scenario: Verify default binding
    Given HTTP server "stub-service"
    And HTTP server timeout is 600000 ms
    Then expect HTTP request body: default
    And receive POST /default

 Scenario: Remove resources
    Given delete Kubernetes service stub-service
    Given delete Pipe binding
