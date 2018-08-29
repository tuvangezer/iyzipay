# iyzipay
## Unofficial Go wrapper for iyzipay API
Every api endpoint covered in the test below is working as intended. For use examples check out how they are done in the file: iyzipay_test.go
### Tests(as of 29-Aug-2018):
```
=== RUN   TestOne
--- PASS: TestOne (0.00s)
=== RUN   TestCreateCard
--- PASS: TestCreateCard (0.38s)
=== RUN   TestAddCard
--- PASS: TestAddCard (1.39s)
=== RUN   TestDeleteCard
--- PASS: TestDeleteCard (0.41s)
=== RUN   TestGetCard
--- PASS: TestGetCard (1.87s)
=== RUN   TestInstallmentInformation
--- PASS: TestInstallmentInformation (0.77s)
=== RUN   TestPayment
--- PASS: TestPayment (0.30s)
=== RUN   TestThreeDSInit
--- PASS: TestThreeDSInit (0.44s)
=== RUN   TestGetPayment
--- PASS: TestGetPayment (0.35s)
=== RUN   TestCancelPayment
--- PASS: TestCancelPayment (0.42s)
=== RUN   TestBKMInit
--- PASS: TestBKMInit (0.91s)
=== RUN   TestBKMGet
--- PASS: TestBKMGet (1.09s)
=== RUN   TestRefundPayment
--- PASS: TestRefundPayment (0.35s)
=== RUN   TestCheckoutFormInitialize
--- PASS: TestCheckoutFormInitialize (0.11s)
=== RUN   TestCheckoutForm
--- PASS: TestCheckoutForm (0.14s)
PASS
```
