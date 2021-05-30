# JAMS Client for Go

A Go client for the JAMS scheduler and (later) a set of useful tools built on it.It doesn't map all possible API's available, but who cares about the Theme API?

If you care about the Theme API (or the other skipped API's) create an issue and explain why they're so useful.

### TODO: 

1. The GetJobLog function doesn't return the log like it should. 
2. The History API has a endpoint that will take oData queries, but, I didn't see a good odata query client for Go. One would assume someone had already done the work to make an odata client to generate the query string, but I can't find one. Maybe I'll build one...
3. This is built against the 6.X release of the JAMS Rest API. I plan on working on the 7.X version soon.


### Planned Tools:

* jamscli - A cli that can interact with the JAMS server
* heidibot - A daemon that monitors for failing jobs and can notify on things
like Microsoft Teams, Slack, or email.
* watcher - Notifies when jobs start or stop using OS notifications

