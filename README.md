# hk_wc - A Golang based tool that mimics unix's `wc`

## Setup
- Execute `make build` to execute the build for this project. This will create a binary in the location `/usr/local/bin`.

  *Note: If make fails due to permission error, try `sudo make build`*
- If not already added, add `/usr/local/bin` to `PATH`.
- To clear the binary, go to the project folder and execute `make clear`

  *Note: If make fails due to permission error, try `sudo make clear`*

## Usage
### `-c`: Count the number of bytes in the input file
<img width="337" alt="image" src="https://github.com/dsvinod90/hk_wc/assets/26185142/5b2784c8-53f9-4910-959b-f0656583b01a">

### `-l`: Count the number of lines in the input file
<img width="346" alt="image" src="https://github.com/dsvinod90/hk_wc/assets/26185142/9a8e210d-a46d-4399-bf9c-6aca083be692">

### `-w`: Count the number of words in the input file
<img width="339" alt="image" src="https://github.com/dsvinod90/hk_wc/assets/26185142/5f86dc3e-1da4-44ca-b35c-fc3f54f574eb">

### `-m`: Count the number of characters in the input file
<img width="343" alt="image" src="https://github.com/dsvinod90/hk_wc/assets/26185142/a6f03b3a-7d60-4a06-b5c0-56c86eb76b80">

### When no flags are given, prints all the statistics for the input file
<img width="321" alt="image" src="https://github.com/dsvinod90/hk_wc/assets/26185142/a4ec2798-c317-43fc-8352-3b0273b63f64">

### Using `STDIN` instead of file input
<img width="416" alt="image" src="https://github.com/dsvinod90/hk_wc/assets/26185142/1828e21d-2d86-4082-854e-860a7fdf34bb">





