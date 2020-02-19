import 'dart:async';

import 'package:flutter/services.dart';

class GoImage {
  static const MethodChannel _channel =
      const MethodChannel('goImage');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }
}
