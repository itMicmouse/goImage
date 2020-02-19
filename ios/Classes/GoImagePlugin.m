#import "GoImagePlugin.h"
#if __has_include(<goImage/goImage-Swift.h>)
#import <goImage/goImage-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "goImage-Swift.h"
#endif

@implementation GoImagePlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftGoImagePlugin registerWithRegistrar:registrar];
}
@end
