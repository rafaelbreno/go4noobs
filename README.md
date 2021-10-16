# Go
<div style="text-align:center">
    <img src="https://github.com/rafaelbreno/go4noobs/blob/master/.src/RAMISTI.png?raw=true" alt="Go4Noobs" />
</div>
<small>Art by <a href="https://twitter.com/loraoraora_">@Lora</a></small>

-------

- __Disclaimer!!__
- __ 👨‍💻 👩‍💻 In development 🚧 🛠 __
- This repo __must__ and __will__ suffer some __major__ updates and modifications along with my learning curve
- Feel totally free to fork and PR this repo
- Enjoy!

## Thank You
- A special thanks to *Ellen Körbes*, for bringing an amazing GOLang content to Brazil in PT-BR
    - [Twitter](https://twitter.com/ellenkorbes) || [GOLang Playlist<small>(in pt-br)</small>](https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)

## Summary
0. [>>Introduction<<](./00_introduction)
1. [Variables and Types](./01_variables_and_types)
    01. [_Hello World_](./01_variables_and_types/01_Hello_World)
    02. [_Println()_](./01_variables_and_types/02_Println)
    03. [_Gopher_](./01_variables_and_types/03_gopher)
    04. [_Scope_](./01_variables_and_types/04_Scope)
    05. [_Types_](./01_variables_and_types/05_types)
    06. [_Custom Type_](./01_variables_and_types/06_custom_type)
    07. [_Type Conversion_](./01_variables_and_types/07_conversion)
    08. [_Exercises_](./01_variables_and_types/08_exercises)
2. [Programming Foundations I](./02_programming_foundations)
    01. [_Boolean_](./02_programming_foundations/01_boolean)
    02. [_String_](./02_programming_foundations/02_string)
    03. [_Constants_](./02_programming_foundations/03_const)
    04. [_Iota_](./02_programming_foundations/04_iota)
    05. [_Part 1 - Exercises_](./02_programming_foundations/05_01_exercises)
    06. [_Procedural Programming_](./02_programming_foundations/06_procedural)
    07. [Loop _For_](./02_programming_foundations/07_loop_for)
    08. [_Conditional_](./02_programming_foundations/08_conditionals)
    09. [_Logical Operator_](./02_programming_foundations/09_logical_operators)
3. [Data Structures I](./03_data_structures)
    01. [_Array_](./03_data_structures/01_array)
    02. [_Slices_](./03_data_structures/02_slice)
    03. [_Slice of Slices_](./03_data_structures/03_slice_of_slice)
    04. [_Slice Make_](./03_data_structures/04_slice_make)
    05. [_Maps_](./03_data_structures/05_maps)
    06. [_Exercises_](./03_data_structures/06_exercises)
    07. [_Structs_](./03_data_structures/07_struct)
4. [Programming Foundations II](./04_programming_foundations_2)
    1. [_Simple Function_](./04_programming_foundations_2/01_simple_function)
    02. [_Variadic Functions_](./04_programming_foundations_2/02_variadic_functions)
    03. [_Multiple Return_](./04_programming_foundations_2/03_multiple_return)
    04. [_Defer Statement_](./04_programming_foundations_2/04_defer_statement)
    05. [_Methods_](./04_programming_foundations_2/05_methods)
    06. [_Interfaces_](./04_programming_foundations_2/06_interfaces)
    07. [_Literal Functions_](./04_programming_foundations_2/07_literal_funcs)
    08. [_Function as Expression_](./04_programming_foundations_2/08_func_as_exp)
    09. [_Return Function_](./04_programming_foundations_2/09_return_func)
    10. [_Callbacks_](./04_programming_foundations_2/10_callbacks)
    11. [_Closure_](./04_programming_foundations_2/11_closure)
    12. [_Recursive_](./04_programming_foundations_2/12_recursive)
    12. [_Pointer_](./04_programming_foundations_2/13_pointer)
5. [Data Structures II](./05_data_structures_2)
    1. [Struct Tags](./05_data_structures_2/01_struct_tags)
    2. [Channels](./05_data_structures_2/02_channels)
6. [Concurrency](./06_concurrency)
    1. [Go Routines](./06_concurrency/01_go_routines)
    2. [Channels](./06_concurrency/02_channels)
    3. [Channel of Channel](./06_concurrency/03_channel_of_channel)
    4. [Parallelization](./06_concurrency/04_parallelization)
    5. [Leaky Buffer](./06_concurrency/05_leaky_buffer)
7. [Packages](./07_packages)
    1. [JSON](./07_packages/01_json)
    2. [Go Template](./07_packages/02_go_template)
99. _Projects_ 
    1. [Go Request Validator](https://github.com/rafaelbreno/go-request-validator)
    2. [Go Twitch Bot](https://github.com/rafaelbreno/go-bot)

## Installing GO in Ubuntu
- Download package from the [Official Download Page](https://golang.org/dl/)
    - In my case the filename is _"go1.15.2.linux-amd64.tar.gz"_
- Open your terminal, and follow these steps
- > `$ mkdir ~/go`
- > `$ sudo mv ~/Downloads/go1.15.2.linux-amd64.tar.gz ~/go`
- > `$ cd ~/go`
- > `$ sudo tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz`
- Copy these:
    -   ```shell
            export GOPATH=$HOME/go
            export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
        ```
- Profile installation
    - Paste Here
        - > `$ sudo nano ~/.profile`
        - > `$ /bin/bash -c ‘source ~/.profile'`
- System-wide installation
    - > `$ sudo nano /etc/profile`
    - > `$ /bin/bash -c ‘source /etc/profile’`
- Re-open your terminal
- Test with:
    - `$ go version`
