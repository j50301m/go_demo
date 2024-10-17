# Kgs Error

## Usage

Create a KgsErr

``` go
func foo() *kgserr.KgsError {
    // Case 1.
   kgsErr := kgserr.New(kgserr.InternalServerError, "Your error message")

   // Case 2: New a kgserr with other error
   otherErr := someService()
   kgsErr = kgserr.New(kgserr.InternalServerError, "Your error message",otherErr)

   return kgsErr
}
```

Compare with kgsCode

``` go
func foo() {
    kgsErr := kgserr.New(kgserr.InternalServerError, "Your error message")
    if kgsErr.Code() == kgserr.InternalServerError {
        // Do something...
    }
}
```

Log a kgsError 

``` go
func foo(ctx context.Context) {
    kgsErr := kgserr.New(kgserr.InternalServerError, "Your error message")

    // If you want some field to log also, use `NewField()` to record with your err
    kgsotel.Error(ctx, err.Message(), kgsotel.NewField("token", req.AccessToken))

    // Or you can just send log simply
    kgsotel.Error(ctx, err.Error())
}
```
