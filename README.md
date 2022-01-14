# Firebase Auth Plugin

Firebase authentication token verification plugin for [Riposo](https://github.com/riposo/riposo).

## Configuration

The following additional configuration options can be used:

| Option                           | Type       | Description                                | Default |
| -------------------------------- | ---------- | ------------------------------------------ | ------- |
| `auth.firebase.project_id`       | `string`   | The Firebase Project ID                    | _none_  |
| `auth.firebase.only_domains`     | `string[]` | Restrict to the following domains          | _none_  |
| `auth.firebase.allow_unverified` | `bool`     | Allow user with unverified email addresses | `false` |

## License

Copyright 2021-2022 Black Square Media Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this material except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
