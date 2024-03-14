## General constraints

- Write on file volumes are always sync writing on data file and index as well
- There just one master sever
- A volume server is registered to a master server through a periodically heartbeat
- It may support local volumes and remote ones
- In case of failure and the index or data file corrupt the entries that are corrupted are discarded
- Each write will replicate to other volumes 
- Deleting a file means to mark them as deleted and posteriorly is reclaimed and swapped 
- Volume size is set at volume server start; this value will be used to know how many volumes it can set 
- Master server will assign a `vid` to a file that will be used across multiple volume server to redundancy
- Each volume server will had a fixed volume size and a fixed volume count