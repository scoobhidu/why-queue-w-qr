import 'package:flutter/material.dart';

/// Checks if string is phone number
bool isValidPhone(String? inputString, {bool isRequired = false}) {
  bool isInputStringValid = false;

  if ((inputString == null ? true : inputString.isEmpty) && !isRequired) {
    isInputStringValid = true;
  }

  if (inputString != null) {
    if (inputString.length > 16 || inputString.length < 6) return false;

    const pattern = r'^[6-9]{1}[0-9]{9}$';

    final regExp = RegExp(pattern);

    isInputStringValid = regExp.hasMatch(inputString);
  }

  return isInputStringValid;
}

/// Checks if string is email.
bool isValidEmail(String? inputString, {bool isRequired = false}) {
  bool isInputStringValid = false;

  if ((inputString == null ? true : inputString.isEmpty) && !isRequired) {
    isInputStringValid = true;
  }

  if (inputString != null) {
    const pattern =
        r'^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$';

    final regExp = RegExp(pattern);

    isInputStringValid = regExp.hasMatch(inputString);
  }

  return isInputStringValid;
}

bool isValidPAN(String? inputString, {bool isRequired = false}) {
  bool isInputStringValid = false;
  if ((inputString == null ? true : inputString.isEmpty) && !isRequired) {
    isInputStringValid = true;
  }
  if (inputString != null) {
    const pattern = r'^[A-Z]{5}\d{4}[A-Z]$';
    final regExp = RegExp(pattern);
    isInputStringValid = regExp.hasMatch(inputString);
  }
  return isInputStringValid;
}

bool isValidAadhar(String? inputString, {bool isRequired = false}) {
  bool isInputStringValid = false;
  if ((inputString == null ? true : inputString.isEmpty) && !isRequired) {
    isInputStringValid = true;
  }
  if (inputString != null) {
    const pattern = r'[0-9]{12}$';
    final regExp = RegExp(pattern);
    isInputStringValid = regExp.hasMatch(inputString);
  }
  return isInputStringValid;
}

bool isValidName(String? inputString, {bool isRequired = false}) {
  bool isInputStringValid = false;
  if ((inputString == null ? true : inputString.isEmpty) && !isRequired) {
    isInputStringValid = true;
  }
  if (inputString != null) {
    const pattern = r'^[A-Za-z]+[A-Za-z ]+$';
    final regExp = RegExp(pattern);
    isInputStringValid = regExp.hasMatch(inputString);
  }
  return isInputStringValid;
}

bool isValidAddress(String? inputString, {bool isRequired = false}) {
  bool isInputStringValid = false;
  if ((inputString == null ? true : inputString.isEmpty) && !isRequired) {
    isInputStringValid = true;
  }
  if (inputString != null) {
    const pattern = r'^[a-zA-z0-9_-\s+]{1,200}$';
    final regExp = RegExp(pattern);
    isInputStringValid = regExp.hasMatch(inputString);
  }
  return isInputStringValid;
}

bool isValidZipCode(String? inputString, {bool isRequired = false}) {
  bool isInputStringValid = false;
  if ((inputString == null ? true : inputString.isEmpty) && !isRequired) {
    isInputStringValid = true;
  }
  if (inputString != null) {
    const pattern = r'[1-9]{1}[0-9]{5}$';
    final regExp = RegExp(pattern);
    isInputStringValid = regExp.hasMatch(inputString);
  }
  return isInputStringValid;
}

// String? validatePAN(
//     String? value, TextEditingController controller, Rx<bool> checkValidator) {
//   print(value);
//   if (controller.text.length == 10 && isValidPAN(value, isRequired: true)) {
//     checkValidator.value = true;
//     print(checkValidator);
//   } else {
//     checkValidator.value = false;
//     print(checkValidator);
//     if (controller.text.length > 1)
//       return 'Please Enter Valid PAN No.';
//     else
//       return null;
//   }
//   return null;
// }
//
// String? validateName(
//     String? value, TextEditingController controller, Rx<bool> checkValidator) {
//   print(value);
//   if (isValidName(value, isRequired: true)) {
//     checkValidator.value = true;
//     print(checkValidator);
//   } else {
//     checkValidator.value = false;
//     print(checkValidator);
//     return 'Please Enter Valid Name';
//   }
//   return null;
// }
//
// String? validatePANorAadhar(
//     String? value, TextEditingController controller, Rx<bool> checkValidator) {
//   print(value);
//   if (isValidPAN(value, isRequired: true) ||
//       isValidAadhar(value, isRequired: true)) {
//     checkValidator.value = true;
//     print(checkValidator);
//   } else {
//     checkValidator.value = false;
//     print(checkValidator);
//     return 'Please Enter Valid PAN or Aadhar';
//   }
//   return null;
// }
//
// String? validateAddress(
//     String? value, TextEditingController controller, Rx<bool> checkValidator) {
//   print(value);
//   if (isValidAddress(value, isRequired: true)) {
//     checkValidator.value = true;
//     print(checkValidator);
//   } else {
//     checkValidator.value = false;
//     print(checkValidator);
//     return 'Please enter valid Address';
//   }
//   return null;
// }
//
// String? validateZipCode(
//     String? value, TextEditingController controller, Rx<bool> checkValidator) {
//   print(value);
//   if (isValidZipCode(value, isRequired: true)) {
//     checkValidator.value = true;
//     print(checkValidator);
//   } else {
//     checkValidator.value = false;
//     print(checkValidator);
//     return 'Please enter valid ZipCode';
//   }
//   return null;
// }
