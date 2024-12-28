namespace go dcdn_seed


enum ErrorCode{
    Success = 0,

    GetPeerFileListFailed = -1,
    AddPeerSeedFailed = -2,
    GetPeerInfoFailed  = -3,
    AddPeerInfoFailed = -4,
    URLParseFailed = -5,
    GetTaskFailed = -6
}

struct FileInfo {
    1: required i32 AppID (api.body="app_id")
    2: required string Fid (api.body="fid")
    3: optional string URL (api.body="url")
}

struct ReportSeedReq {
    1: required string PeerID (api.body="peer_id")
    2: required string ServiceAddress (api.body="service_addr")
    3: required list<FileInfo> Files (api.body="files")
}

struct BaseResponse {
    1: required i32 Code (api.body="code")
    2: required string Message (api.body="message")
}

struct DeviceBasicQueryReq {
    1: required string PeerID (api.query="peer_id")
}

struct DownloadTaskInfo {
    1: required list<FileInfo> Tasks(api.body="tasks")
}

struct QueryFidPeerReq {
    1: required i32 AppID (api.query="app_id")
    2: required string URL (api.query="url")
    3: optional string FID (api.query="fid")
}

struct PeerInfo {
    1: required string PeerID (api.body="peer_id")
    2: required string ServiceAddress (api.body="service_addr")
}

struct QueryFidPeerData {
    1: required list<PeerInfo> PeerList (api.body="peer_list")
}

service SeedService {
    BaseResponse  ReportSeedAll(1: ReportSeedReq request) (api.post="/seed_manager/device/report/all")
    DownloadTaskInfo GetDownloadTask(1: DeviceBasicQueryReq request) (api.get="/seed_manager/device/task")

    QueryFidPeerData QueryFidPeer(1: QueryFidPeerReq request)  (api.get="/seed_manager/sdk/list")
}