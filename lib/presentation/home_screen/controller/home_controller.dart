import 'package:flutter/material.dart';
import 'package:get/get_rx/src/rx_types/rx_types.dart';
import 'package:get/get_state_manager/src/simple/get_controllers.dart';

import '../models/home_model.dart';

class HomeController extends GetxController{

  HomeModel stockDetailsModelObj = HomeModel();
  TextEditingController usernameController = TextEditingController();
  FocusNode usernameNode = FocusNode(
    canRequestFocus: true,
  );
  @override
  void onInit() async {
    super.onInit();
  }

  @override
  void onReady() {
    super.onReady();
  }

  @override
  void onClose() {
    super.onClose();
  }
}
