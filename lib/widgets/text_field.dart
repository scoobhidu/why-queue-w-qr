import 'package:flutter/material.dart';
import 'package:why_queue_w_qr/core/utils/color_constant.dart';

import '../core/utils/math_utils.dart';
import '../theme/app_style.dart';

class CustomTextField extends StatelessWidget {
  CustomTextField({
    required this.hintText,
    required this.controller,
    required this.margin,
    required this.node,
    required this.keyboardType,
    this.textStyle = const TextStyle(
      color: LightTheme.bluegray900,
      fontWeight: FontWeight.w400,
      fontFamily: 'Sora',
      fontSize: 14,
    ),
    this.containerDecoration = const BoxDecoration(
      color: LightTheme.whiteA700,
      borderRadius: BorderRadius.all(Radius.circular(9)),
    ),
  });

  String hintText;
  TextEditingController controller;
  EdgeInsets margin;
  FocusNode node;
  TextInputType keyboardType;
  BoxDecoration containerDecoration;
  TextStyle textStyle;

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: margin,
      decoration: containerDecoration,
      child: TextField(
        controller: controller,
        autofocus: true,
        textInputAction: TextInputAction.done,
        style: textStyle,
        strutStyle: StrutStyle.disabled,
        keyboardType: keyboardType,
        focusNode: node,
        decoration: InputDecoration(
          border: InputBorder.none,
          hintText: hintText,
          contentPadding: EdgeInsets.symmetric(horizontal: 15),
        ),
        onTap: (){
          node.requestFocus();
        },
      ),
    );
  }
}
