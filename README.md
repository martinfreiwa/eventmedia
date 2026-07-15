# eventmedia

`eventmedia` is a small, dependency-free Go package for validating operational plans for guest photo, video, and voice-message collection.

```go
result := eventmedia.Validate(plan, time.Now())
if !result.Ready {
    // Review result.Findings before the event opens.
}
```

It checks HTTPS destinations, privacy wording, removal contacts, moderation, live-display safety, retention, upload windows, network tests, fallback plans, and responsible roles.

The library is platform-neutral. For an implementation example of a browser-based QR upload workflow, see [Gathmo's event media flow](https://gathmo.com/how-it-works).

## Install

```bash
go get github.com/martinfreiwa/eventmedia@v0.1.0
```

## Test

```bash
go test ./...
```

## License

MIT
