import 'dart:async';
import 'dart:developer';

import 'package:camera/camera.dart';
import 'package:device_info_plus/device_info_plus.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:get/get.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:themed/themed.dart';
import 'package:why_queue_w_qr/routes/app_routes.dart';
import 'core/utils/color_constant.dart';
import 'core/utils/pref_utils.dart';

late List<CameraDescription> cameras;

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  cameras = await availableCameras();

  final deviceInfoPlugin = DeviceInfoPlugin();
  final deviceInfo = await deviceInfoPlugin.deviceInfo;
  final map = deviceInfo.toMap();
  print(map);
  print(map["id"]);
  print(map["brand"]);

  await Get.put(PrefUtils());

  SystemChrome.setSystemUIOverlayStyle(SystemUiOverlayStyle(
      statusBarIconBrightness: Brightness.light, // dark text for status bar
      statusBarColor: Colors.transparent));
  SystemChrome.setPreferredOrientations([DeviceOrientation.portraitUp])
      .then((value) => runApp(MyApp()));
}

class MyApp extends StatefulWidget {
  @override
  State<MyApp> createState() => _MyAppState();
  var netInfo;
  Rx<bool> connected = true.obs;
}

class _MyAppState extends State<MyApp> with WidgetsBindingObserver {
  // This widget is the root of your application.

  @override
  void initState() {
    WidgetsBinding.instance?.addObserver(this);
    widget.netInfo = InternetConnectionChecker().onStatusChange.listen((status) {
      switch(status) {
        case InternetConnectionStatus.connected:
          log("Connected to net");
          widget.connected.value = true;
          break;
        default:
          log("Not connected anymore");
          widget.connected.value = false;
      }
    });
    super.initState();
  }

  @override
  void dispose() {
    WidgetsBinding.instance?.removeObserver(this);
    super.dispose();
  }

  @override
  void didChangePlatformBrightness() {
    var brightness = Theme.of(context).brightness;
    print(brightness);
    if (brightness == Brightness.dark) {
      Themed.currentTheme = DarkTheme;
    } else {
      Themed.clearCurrentTheme();
    }
    super.didChangePlatformBrightness();
  }

  @override
  Widget build(BuildContext context) {
    return Obx(() => AbsorbPointer(
      absorbing: widget.connected.value,
      // absorbing will depend on an active internet connection
      child: Themed(
        child: GetMaterialApp(
          debugShowCheckedModeBanner: false,
          locale: Get.deviceLocale,
          //for setting localization strings
          fallbackLocale: Locale('en', 'US'),
          title: 'Why Queue with QR',
          // initialBinding: InitialBindings(),
          initialRoute: AppRoutes.initialRoute,
          getPages: AppRoutes.pages,
          builder: (context, child) => MediaQuery(
              data: MediaQuery.of(context).copyWith(
                  textScaleFactor:
                  MediaQuery.of(context).textScaleFactor.clamp(1, 1.2)),
              child: child!),
          // TODO: will set theme in next commit
          // theme: ThemeData(
          //     colorScheme: Theme.of(context).colorScheme.copyWith(
          //         primary: LightTheme.deepOrange400,
          //         secondary: LightTheme.bluegray900))
        ),
      ),
    ));
  }
}
