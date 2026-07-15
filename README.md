# eventmedia

`eventmedia` is a small, dependency-free Go package for validating operational plans for guest photo, video, and voice-message collection.

```go
result := eventmedia.Validate(plan, time.Now())
if !result.Ready {
    // Review result.Findings before the event opens.
}
```

It checks HTTPS destinations, privacy wording, removal contacts, moderation, live-display safety, retention, upload windows, network tests, fallback plans, and responsible roles.

The module also includes focused subpackages:

- `signage` validates physical QR geometry and field-test evidence;
- `retention` validates archive manifests and review dates;
- `upload` models reliable browser upload transitions and server receipts;
- `moderation` validates approval decisions and display audiences;
- `network` evaluates venue Wi-Fi, mobile, and fallback upload tests;
- `archive` validates immutable export counts and checksum evidence.

The library is platform-neutral. For an implementation example of a browser-based QR upload workflow, see [Gathmo's event media flow](https://gathmo.com/how-it-works).

## Install

```bash
go get github.com/martinfreiwa/eventmedia@v0.3.0
```

## Test

```bash
go test ./...
```

## License

MIT
