# JAMS Client for Go

A Go client for the JAMS scheduler and (later) a set of useful tools built on it.It doesn't map all possible API's available, but who cares about the Theme API?

If you care about the Theme API (or the other skipped API's) create an issue and explain why they're so useful.

### Examples

* submit - Find and submit a job
* schedule - Look at Jobs and Setups to make a simplified weekly job schedule

### TODO: 

1. The GetJobLog function doesn't return large logs.
2. This is built against the 6.X release of the JAMS Rest API. I plan on working on the 7.X version at some point.


### Planned Tools:

* jamscli - A cli that can interact with the JAMS server
* heidibot - A daemon that monitors for failing jobs and can notify on things
like Microsoft Teams, Slack, or email.
* watcher - Notifies when jobs start or stop using OS notifications

