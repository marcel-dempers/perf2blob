# perf2blob

Work in Progress:

Tool to run Linux Perf and Push the results to an Azure blob storage

Wrapper around `linux-tools` perf utility.
Runs perf and writes `perf.data` to Azure Blob. 

Use Case: Can be used on a container environment to run perf on the host and push the data to blob for external analysis (Like Generate flame graphs)

## Usage

Record a process for 30 sec

```
export AZURE_STORAGE_ACCOUNT_NAME=test
export AZURE_STORAGE_ACCOUNT_KEY=XXXX
export AZURE_STORAGE_CONTAINER=test

./perf2blob record -p $PROCESS_ID -ag -F 97 -- sleep 30
```

### Docker 

```
#!/bin/bash
docker run -it \
--rm \
--privileged \
--ipc host \
--pid host \
-v $PWD:/out \
-v /var/lib/docker:/var/lib/docker \
-v /sys:/sys:ro \
-v /etc/localtime:/etc/localtime:ro \
-v /run:/run \
-v /var/log:/var/log \
--net host \
aimvector/perf2blob record -p 1 -ag -F 97 -o /out/perf.datak -- sleep 5
```