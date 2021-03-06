**GO-WebRados**
---------

**Go-WebRados** is a simple and high performance HTTP service for **CEPH** distributed file system. 
The main goal of this tool is to provide simple HTTP interface for **Ceph's** bare **RADOS** layer.
**Go-WebRados** is not a replacement for **RadosGW** as is does not have all reach APIs and features of RadosGW(S3, Swift, etc ...), instead it stands for simplicity and performance. 

GO-WebRados relies on HTTP methods to interact with RADOS object, thus it can provide access to objects store in RADOS directly from internet browser . 
The ide is to have web accessible storage for millions of relatively small files, which can be accessed from browser directly. 

GO-WebRados relies on C bindings of **Ceph** so in order to run this program you need to install Ceph packages. 
Running Ceph services on computer which hosts GO-WebRados is not required, it's even better to have a dedicated server or server for running GO-WebRados  

### **Download and install**
---------

You can build GO-WebRados from source or download precompiled binaries. If you already have installed Cephs packages and want to make things easy , 
just download te GO-WebRados binary, make it executeable, and you are ready to run .  

Building from a source is also easy. 

```shell
git clone  https://github.com/sadoyan/go-webrados.git
cd go-webrados
export GOROOT=/path/to/your/go
./build.sh

./webrados /path/to/config.ini
```



### **Configuration**
---------

Configuration paramaters are stored in ```config.ini``` file, which should be in running directory.
Sample config file, with reasoneable defaulr ships with source code. 

```ini
[main]
listen 0.0.0.0:8080
dispatchers 20
serveruser admin
serverpass 261a5983599fd57a016122ec85599ec4
uploadmaxpart 52428800
dangerzone yes
readonly no
authread no
authwrite yes
radoconns 20
logfile yes
logpath: /opt/webrados.log

[monitoring]
enabled true
url:  127.0.0.1:9090
user: admin
pass: admin
```
### **API**
---------

| **Name** | **Description** |
| ------------- | ------------- |
|**Read File**|HTTP ```GET:``` http://{BINDADDRESS}/{POOLNAME}/{FILENAME}|
|**Read File**|HTTP ```GET:``` http://{BINDADDRESS}/{POOLNAME}/{FILENAME}|
|**Upload File**|HTTP ```POST, PUT:``` http://{BINDADDRESS}/{POOLNAME}/{FILENAME}|
|**Remove File**|HTTP ```DELETE:``` http://{BINDADDRESS}/{POOLNAME}/{FILENAME}|

Configuration file is pretty simple and intuitive. 

### **Section main**
---------

| **Name**  | **Description** |
| ------------- | ------------- |
|**listen**|IP port to bind.|
|**dispatchers**|Number of threads for webserver.|
|**serveruser**|Static user.|
|**serverpass**|MD5 hash of password for static user. It can be the output of `echo -n SecretPaSs | md5sum |awk '{print $1}'`|
|**dangerzone**|Enable destructive methods and commands (DELETE).|
|**readonly**|Enable readonly mode. If 'yes' only GET is allowed.|
|**authread**|Require authentication for GET only.|
|**authwrite**|Require authentication for POST/PUT/DELETE.|
|**radoconns**|Number of connection to CEPH.|
|**logfile**|Log to file, if 'no' logs are sent to stdout.|
|**logpath**|Path for log file.|
|**uploadmaxpart**| Maximum file chunk size (Sbould be amaller or erqual ro `osd max object size`).| 
|**allpools:**|yes/no . If yes program will scan ceph and enable access via web to all pool.| 
|**poollist:**|Works only if **allpools** is set to **no**. Coma separated list of pools which should be accesible via webrados program.|  
|**redis**|IP address and port of Redis server.|
### **Section monitoring**
---------
| **Name**  | **Description** |
| ------------- | ------------- |
|**enabled**|Enable/Disable monitoring.|
|**url**|IP address and port for minitoring interface.|
|**user**|Monitoring user.|
|**pass**|Password for monitoring user.|

### **users.txt file**

GO-Webrados can dynamically update users from ```users.txt``` file . 
```users.txt``` should contain user and md5hash of password divided by space in each line.  
`echo -n SecretPaSs | md5sum |awk '{print $1}'` on Linux systemd will output md5hash for using it as password in `users.txt` file 
GO-Webrados will periodically read ```uesrs.txt``` file and automatically update users in memory. 

### **Large files**

In order to be able to store large file in RADOS directly files needs to be split to smaller chunks. 
GO-WebRados will automatically set maximum chunk size to  **uploadmaxpart** and split files in accordance to that. 
For this GO-WebRados requires Redis servers to store metadata for big files. (Make sure to keep Redis database safe and periodically back it up.)
Redis will store metadata only for large files. Small files are to split and not require special metadata server.  

### **Special commands**

**HTTP GET** http://{BINDADDRESS}/{POOLNAME}/{FILENAME}?info 
Return information about requested file in json format. 

```curl -s  http://ceph1:8080/bublics/katana.mp4?info | python -mjson.tool```

```json
{
    "name": "katana.mp4",
    "pool": "bublics",
    "segments": "11",
    "size": "471861144"
}
```
