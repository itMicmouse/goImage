import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:goImage/goImage.dart';

void main() {
  const MethodChannel channel = MethodChannel('goImage');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await GoImage.platformVersion, '42');
  });
}
