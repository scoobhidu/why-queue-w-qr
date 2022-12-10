import 'package:flutter/material.dart';

import '../core/utils/color_constant.dart';

class CustomElevatedButton extends StatelessWidget {
  CustomElevatedButton({
    required this.buttonColor,
  });

  MaterialStateProperty<Color> buttonColor;

  @override
  Widget build(BuildContext context) {
    return Container(
      child: ElevatedButton(
        onPressed: (){},
        child: Text("Sign In"),
        style: ButtonStyle(
          backgroundColor: buttonColor,
          shadowColor: buttonColor,
          elevation: MaterialStateProperty.all(6),
        ),
      ),
    );
  }
}
