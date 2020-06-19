# OpenWeatherMap Console 
## Go files 
### /main/owm.go
* Displays current weather information for a given US ZIP code supported by OpenWeatherMap.
#### Usage 
The following commands assume the the working directory is /owmconsole/main.
##### Build
````
go build main.go
````
##### Run
````
go run main.go -zip <ZIP code> -extreme <temperature extreme>
````
##### Arguments
````
<ZIP code>
````
* A US ZIP code corresponding to the city of interest. Note that OpenWeatherMap does not provide weather information for all US ZIP codes. 
````
<temperature extreme>
````
* Optional. Display the high temperature for the day with ````high```` and the low temperature for the day with ````low````. 
