namespace RpsTest;

public class RunCode
{
    public static async Task<ResultData> MethodRun(List<string[]> arrays)
    {
        var i = 0;
        var b = 0;

        for (b = 0; b < 50000; b++)
        {
            var array = arrays[b % arrays.Count];
            for (var index = 0; index < array.Length; index++)
            {
                var t = array[index];
                if (t == "qw2")
                    i++;
            }
        }


        await Task.Delay(10);

        for (b = 0; b < 50000; b++)
        {
            var array = arrays[b % arrays.Count];
            for (var index = 0; index < array.Length; index++)
            {
                var t = array[index];
                if (t == "qw5")
                    i++;
            }
        }

        await Task.Delay(10);

        for (b = 0; b < 50000; b++)
        {
            var array = arrays[b % arrays.Count];
            for (var index = 0; index < array.Length; index++)
            {
                var t = array[index];
                if (t == "qw8")
                    i++;
            }
        }

        return new ResultData
        {
            Result = i
        };
    }
}