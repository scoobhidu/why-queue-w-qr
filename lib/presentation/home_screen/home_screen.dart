import 'package:flutter/material.dart';
import 'package:get/get_state_manager/src/simple/get_view.dart';
import 'package:glass_kit/glass_kit.dart';
import 'package:why_queue_w_qr/widgets/custom_button.dart';
import 'package:why_queue_w_qr/widgets/text_field.dart';

import '../../core/utils/color_constant.dart';
import '../../core/utils/math_utils.dart';
import '../../theme/app_style.dart';
import 'controller/home_controller.dart';

class HomeScreen extends GetWidget<HomeController> {
  GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return SafeArea(
        child: Scaffold(
            backgroundColor: LightTheme.whiteA700,
            body: Container(
              decoration: const BoxDecoration(
                image: DecorationImage(
                  image: AssetImage('assets/images/glassmorphAbstract.jpg'),
                  fit: BoxFit.cover,
                )
              ),
              child: GlassContainer.frostedGlass(
                height: size.height,
                width: size.width,
                color: LightTheme.lightlavender,
                child: Column(
                  children: [
                    Text(
                      "Hello Again!",
                      style: AppStyle.txtSoraSemiBold16Bluegray900.copyWith(fontSize: 22),
                    ),
                    Text(
                      "Welcome back you've \n been missed!",
                      style: AppStyle.txtSoraRegular14Bluegray900.copyWith(
                        fontSize: 18,
                        height: 1.3,
                      ),
                      textAlign: TextAlign.center,
                    ),
                    Column(
                      children: [
                        CustomTextField(
                          hintText: "Enrollment Number",
                          controller: controller.usernameController,
                          margin: EdgeInsets.symmetric(horizontal: 40),
                          node: controller.usernameNode,
                          keyboardType: TextInputType.number,
                        ),
                        InkWell(
                          child: Text(
                            "Changed your device?",
                            style: AppStyle.txtSoraRegularExtra12Gray700,
                          ),
                        ),
                      ],
                    ),
                    CustomElevatedButton(
                      buttonColor: MaterialStateProperty.all(LightTheme.cherryRed),
                    ),
                ],
              )
              ),
            )));
  }
}
