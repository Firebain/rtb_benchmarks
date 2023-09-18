namespace RpsTest;

public class RunCode
{
    public static async Task<ResultData> MethodRun(List<List<string>> arrays)
    {
        var i = 0;
        var b = 0;

        for (b = 0; b < 50000; b++)
        {
            var array = arrays[b % arrays.Count];
            for (var index = 0; index < array.Count; index++)
            {
                var t = array[index];
                if (t == "qw2") {
                    i++;

                    break;
                }
            }
        }

        await Task.Delay(10);

        for (b = 0; b < 50000; b++)
        {
            var array = arrays[b % arrays.Count];
            for (var index = 0; index < array.Count; index++)
            {
                var t = array[index];
                if (t == "qw5") {
                    i++;

                    break;
                }
            }
        }

        await Task.Delay(10);

        for (b = 0; b < 50000; b++)
        {
            var array = arrays[b % arrays.Count];
            for (var index = 0; index < array.Count; index++)
            {
                var t = array[index];
                if (t == "qw8") {
                    i++;

                    break;
                }
            }
        }

        return new ResultData
        {
            Result = i
        };
    }
}