# Learn to Code with Go Programming Language
> Course: https://www.udemy.com/course/learn-how-to-code/
>
> Docs for course: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit

## Section 3
**Go Workspace**
- one folder - any name, any location
  - `/bin`: executable
  - `/pkg`: packages from go
  - `/src`: source code that you create

What is cool about this is that we can be anywhere on our machine and just
add the three directories listed 👆🏾, and we have a workspace that will act
just like the old school way of setting up a go project where we have to be at `~/go`

**Go Commands**
- `go version`: print the version
- `go env`: print environment variables
- `go help`: go help command
- `go fmt`: go fmt command which formats your code
    - `./`
- `go run`: is `go build` and `go run` which builds and executes the code in the project


**Go Modules**
This is the default behavior as version 1.13 of Golang

**Go Packages**: think of this as a way to organize code like we do on computer where we have directories an files
## How computers work
- computers run on electricity
- electricity has 2 discrete states:
  - on
  - off
- coding schemes
  - Porch Light
    - If light is on, come trick or treat
    - If light is off, do not come trick or treat
  - 1 porch light = 2 messages
  - 2 porch lights = 4 messages
  - 3 porch lights = 8 messages

The way you figure out how many messages is a mathematical formula where $2^n$ where $n$ is the number of things you have. So if you 4 porch lights that $2^4$ = $16$ messages.

## Measuring Bits
- 1 bit
- 8 bits = **1 byte**
- 1000 bytes = **kilobytes**
- 1000 kilobytes = **megabytes**
- 1000 megabytes = **gigabyte**
- 1000 gigabyte = **terabyte**

One of the first computers played a pivotal role in a world war II. The computer check the trajectory of shells shot from ships. Grace Hooper was on the Navy Computer Scientist noted in her diary literally about a bug that they found in the vacuum tubes of the computer. They called this process **debugging**.

## Numeric Types
Unsigned integer are always positive numbers. While signed integers are either negative or positive.
