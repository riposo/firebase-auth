# Firebase Auth Plugin

Firebase authentication token verification plugin for [Riposo](https://github.com/riposo/riposo).

## Configuration options

- `RIPOSO_AUTH_FIREBASE_TOKEN_URL` - the token verification URL (default: `https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com`)
- `RIPOSO_AUTH_FIREBASE_VERIFY_ISS` - verify `iss` claim, ensure it matches given value
- `RIPOSO_AUTH_FIREBASE_VERIFY_AUD` - verify `aud` claim, ensure it matches given value
- `RIPOSO_AUTH_FIREBASE_ONLY_DOMAINS` - restricts users to given (comma-separated) list of domains
- `RIPOSO_AUTH_FIREBASE_ALLOW_EXPIRED` - allow expired tokens (not recommended)
- `RIPOSO_AUTH_FIREBASE_ALLOW_UNVERIFIED` - allow user with non-verified emails (not recommended)
