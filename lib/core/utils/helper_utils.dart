import 'package:flutter/services.dart';
// import 'package:get/get_core/src/get_main.dart';
import 'package:intl/intl.dart';
import 'package:permission_handler/permission_handler.dart';

class HelperUtils {
  Future<bool> request_CameraPermission() async {
    final Permission permission = Permission.camera;
    PermissionStatus requestStaus = await permission.request();

    bool isPermanetelyDenied = await permission.isPermanentlyDenied;
    // Get.snackbar('', requestStaus.toString());
    // Get.snackbar('', isPermanetelyDenied.toString());

    if (isPermanetelyDenied == true) {
      return false;
    } else {
      return requestStaus.isGranted;
    }
  }

  static bool isPositiveValue(double value) {
    if (value > 0)
      return true;
    else
      return false;
  }

  static String currencyFormat(String amount) {
    var indiaFormat = NumberFormat.currency(
      symbol: '\u{20B9}',
      locale: "en_IN",
      decimalDigits: 2,
    );
    return indiaFormat.format(double.parse(amount));
  }

  static String currencyFormatWithoutDecimal(String amount) {
    var indiaFormat = NumberFormat.currency(
      symbol: '\u{20B9}',
      locale: "en_IN",
      decimalDigits: 0
    );
    return indiaFormat.format(double.parse(amount));
  }

  static String currencyCompactFormat(dynamic amount) {
    return NumberFormat.compactSimpleCurrency(locale: 'en-IN')
        .format(amount)
        .toString()
        .replaceAll('₹', '');
  }
}

class DateTextFieldFormatter extends TextInputFormatter {
  static const _maxChars = 8;
  @override
  TextEditingValue formatEditUpdate(
      TextEditingValue oldValue, TextEditingValue newValue) {
    if (oldValue.text.length >= newValue.text.length) {
      return newValue;
    }
    var dateText = _addSeparator(newValue.text, '/');
    return newValue.copyWith(
        text: dateText, selection: updateCursorPosition(dateText));
  }

  String _addSeparator(String value, String separator) {
    value = value.replaceAll('/', '');
    var newString = '';
    for (int i = 0; i < value.length; i++) {
      newString += value[i];
      if (i == 1) {
        newString += separator;
      }
      if (i == 3) {
        newString += separator;
      }
    }
    return newString;
  }

  TextSelection updateCursorPosition(String text) {
    return TextSelection.fromPosition(TextPosition(offset: text.length));
  }
}
