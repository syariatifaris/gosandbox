# API Sandbox
This API serve as a backend sandbox.

## Dependencies

[Facebook's Grace library](http://github.com/facebookgo/grace) is used for graceful shutdown.

```
go get github.com/facebookgo/grace
```

[Karlkfi Dependency Injection library](http://github.com/karlkfi/inject) is used for golang DI.

```
go get github.com/karlkfi/inject
```

## Usage
1. Sync the go vendor
    ```
    govendor sync -v
    ```

1. Launch the server:

    ```
    go run main.go
    ```

    (ctrl-c to quit)

1. Query Index:

    ```
    $ curl http://localhost:8080/
    ```


## License

   Licensed under the [Apache License Version 2.0](LICENSE) (the "License");
   you may not use this project except in compliance with the License.

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
