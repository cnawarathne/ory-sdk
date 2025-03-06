import 'package:test/test.dart';
import 'package:ory_kratos_client/ory_kratos_client.dart';

// tests for SessionDevice
void main() {
  final instance = SessionDeviceBuilder();
  // TODO add properties to the builder and call build()

  group(SessionDevice, () {
    // A list of authenticators which were used to authenticate the session.
    // BuiltList<SessionAuthenticationMethod> authenticationMethods
    test('to test the property `authenticationMethods`', () async {
      // TODO
    });

    // Device record ID
    // String id
    test('to test the property `id`', () async {
      // TODO
    });

    // IPAddress of the client
    // String ipAddress
    test('to test the property `ipAddress`', () async {
      // TODO
    });

    // Geo Location corresponding to the IP Address
    // String location
    test('to test the property `location`', () async {
      // TODO
    });

    // Is this device trusted? (only matters if this device submitted aal2+ credentials)
    // bool trusted
    test('to test the property `trusted`', () async {
      // TODO
    });

    // UserAgent of the client
    // String userAgent
    test('to test the property `userAgent`', () async {
      // TODO
    });

  });
}
