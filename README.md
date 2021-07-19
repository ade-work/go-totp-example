# Example TOTP usage

_TOTP - Time-based One Time Password_

## Usage 
Use commands in makefile to generate codes. Then send code and time to 
_POST localhost:3000/totp/validate_ 
to verify code.

#### Dependencies:

* [xlzd/gotp](https://github.com/xlzd/gotp)
* [pquerna/otp](https://github.com/pquerna/otp)
* [Fiber](https://github.com/gofiber/fiber)
* [Testify](https://github.com/stretchr/testify)
