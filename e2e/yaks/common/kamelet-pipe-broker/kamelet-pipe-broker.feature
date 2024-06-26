Feature: Camel K can bind Kamelets to the broker

  Background:
    Given Camel K resource polling configuration
      | maxAttempts          | 40   |
      | delayBetweenAttempts | 3000 |

  Scenario: Sending event to the custom broker with Pipe
    Given Camel K integration logger-sink-pipe is running
    Then Camel K integration logger-sink-pipe should print message: Hello Custom Event from sample-broker

  Scenario: Remove resources
    Given delete Camel K integration timer-source-pipe
    Given delete Camel K integration logger-sink-pipe
