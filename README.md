# ais-stream

AIS stream handler to load data to MongoDB

## To do

- convert all context.TODO() calls in MongoDB handler to a context with timeout
- add a small ui - a map showing coverage
- update logging so it shows something a bit more informative than just the de-duplicator
- confirm that all streams automatically restart (i.e. are resiliant and fault tolerant)
