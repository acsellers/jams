# JAMS
A Go client for the JAMS scheduler and (later) a set of useful tools built on it.It doesn't map all possible API's available, but who cares about the Theme API?

If you care about the Theme API (or the other skipped API's) create an issue and explain why they're so useful.

Currently, I didn't see a good odata query client for Go. One would assume someone had already done the work to make an odata client to generate the queries, but I can't find one. Maybe I'll build one...

This is built against the 6.X release of the JAMS Rest API. I plan on working on the 7.X version soon.

Planned Tools:

* jamscli - A cli that can interact with the JAMS server
* heidibot - A daemon that monitors for failing jobs and can notify on things
like Microsoft Teams or Slack.
* watcher - Notifies when a job starts or stops using OS notifications

