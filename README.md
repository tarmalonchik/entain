
## Run

To run docker compose with `Application` and `PG`, you need run this command 

```
docker compose up -d 
```

## Default

By default, application will listen on `8080` port

## Notes

+ `UserMaxBalance` is limited with 92,233,720,368,547,758.07 number, this is because I keep my money in cents and the maximum possible number of cents is 9,223,372,036,854,775,807 
> I decided to keep money in cents, because there were no requirements, regarding the max number for money 
+ Since this was not in the requirements, I decided that the amount of `0.00` money is also a valid amount. So an amount equal to `0.00` will also be processed the same as any other.
