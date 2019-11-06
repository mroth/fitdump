
Quick (very quick) library to parse the exported data from a [Fitbit data
export][1]. Currently this only handles body weight data -- as this is the only
use case I personally had. PRs welcome if you'd like to add more.

Unlike data obtained via the Fitbit API, this contains exact values and
timestamps, rather than interpolated time sequences.

The included example program parses a collection of exported weight data JSON
files into a sorted CSV file, suitable for use with [HealthKit CSV Importer][2].

[1]: https://help.fitbit.com/articles/en_US/Help_article/1133
[2]: https://2017.lionheartsw.com/software/health-csv-importer/
