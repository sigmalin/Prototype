
namespace AssetLoader
{
    public class AddressableDownloadStatus : IDownloadStatus
    {
        public int TotalFiles { get; set; } = 0;

        public int CompleteFiles { get; set; } = 0;

        public long DownloadedBytes { get; set; } = 0u;

        public long FileBytes { get; set; } = 0u;
    }
}
