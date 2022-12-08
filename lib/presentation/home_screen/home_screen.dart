import 'package:flutter/material.dart';
import 'package:get/get_state_manager/src/simple/get_view.dart';

import '../../core/utils/color_constant.dart';
import '../../core/utils/math_utils.dart';
import 'controller/home_controller.dart';

class HomeScreen extends GetWidget<HomeController> {
  GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return SafeArea(
        child: Scaffold(
            backgroundColor: LightTheme.whiteA700,
            body: Container(
                width: size.width,
                child: Form(
                  key: _formKey,
                  autovalidateMode: AutovalidateMode.onUserInteraction,
                  child: Container(
                    height: size.height,
                    width: size.width,
                    color: LightTheme.lightlavender
                  ),
                ))));
  }
}
