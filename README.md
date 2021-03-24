# Firebase Auth Plugin

Firebase authentication token verification plugin for [Riposo](https://github.com/riposo/riposo).

## Configuration options

- `RIPOSO_AUTH_FIREBASE_TOKEN_URL` - the token verification URL (default: `https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com`)
- `RIPOSO_AUTH_FIREBASE_VERIFY_ISS` - verify `iss` claim, ensure it matches given value
- `RIPOSO_AUTH_FIREBASE_VERIFY_AUD` - verify `aud` claim, ensure it matches given value
- `RIPOSO_AUTH_FIREBASE_ONLY_DOMAINS` - restricts users to given (comma-separated) list of domains
- `RIPOSO_AUTH_FIREBASE_ALLOW_EXPIRED` - allow expired tokens (not recommended)
- `RIPOSO_AUTH_FIREBASE_ALLOW_UNVERIFIED` - allow user with non-verified emails (not recommended)

## License

Copyright 2021 Black Square Media Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this material except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
