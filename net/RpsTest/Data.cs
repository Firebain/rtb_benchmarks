using System.Text.Json.Serialization;

namespace RpsTest;

public struct ResultData
{
    [JsonPropertyName("result")] public int Result { get; set; }
}

public struct AppData
{
    public List<List<string>> Arrays { get; set; }

    public static AppData Init()
    {
        var appData = new AppData();
        appData.Arrays = new List<List<string>>();

        for (var i = 0; i < 20; i++)
        {
            if (i % 2 == 0)
            {
                appData.Arrays.Add(new List<string>(){ "qw1", "qw2", "qw3", "qw4", "qw5"});
            }
            else
            {
                appData.Arrays.Add(new List<string>(){ "qw5", "qw4", "qw3", "qw2", "qw1" });
            }
        }

        return appData;
    }
}