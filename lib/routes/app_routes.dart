import 'package:flutter/material.dart';
import 'package:get/get_navigation/src/routes/get_route.dart';
import 'package:why_queue_w_qr/presentation/home_screen/binding/home_binding.dart';
import 'package:why_queue_w_qr/presentation/home_screen/home_screen.dart';
import 'package:why_queue_w_qr/presentation/stock_details_screen/binding/stock_details_binding.dart';
import 'package:why_queue_w_qr/presentation/stock_details_screen/stock_details_screen.dart';

class AppRoutes {
  static String tellUsAboutYourselfMaritalScreen = '/stock_details_screen';
  static String initialRoute = '/home_screen';

  static List<GetPage> pages = [
    GetPage(
      name: tellUsAboutYourselfMaritalScreen,
      page: () => StockDetailsScreen(),
      bindings: [
        StockDetailsBinding(),
      ],
    ),
    GetPage(
      name: initialRoute,
      page: () => HomeScreen(),
      bindings: [
        HomeBindings(),
      ],
    ),
  ];
}
