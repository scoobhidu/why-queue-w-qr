//ignore: unused_import
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

class PrefUtils {
  static SharedPreferences? _sharedPreferences;

  PrefUtils() {
    SharedPreferences.getInstance().then((value) {
      _sharedPreferences = value;
    });
  }

  Future<void> init() async {
    _sharedPreferences ??= await SharedPreferences.getInstance();
    print('SharedPreference Initialized');
  }

  ///will clear all the data stored in preference
  void clearPreferencesData() async {
    _sharedPreferences!.clear();
  }

  Future<void> setsessionid(String value) {
    return _sharedPreferences!.setString('sessionid', value);
  }

  String getsessionid() {
    try {
      return _sharedPreferences!.getString('sessionid') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setssosessionid(String value) {
    return _sharedPreferences!.setString('ssosessionid', value);
  }

  String getssosessionid() {
    try {
      return _sharedPreferences!.getString('ssosessionid') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setphoneno(String value) {
    return _sharedPreferences!.setString('phone_no', value);
  }

  Future<void> setclientname(String value) {
    return _sharedPreferences!.setString('client_name', value);
  }

  String getclientname() {
    try {
      return _sharedPreferences!.getString('client_name') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setclientid(String value) {
    return _sharedPreferences!.setString('client_id', value);
  }

  String getclientid() {
    try {
      return _sharedPreferences!.getString('client_id') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setpan(String value) {
    return _sharedPreferences!.setString('pan', value);
  }

  String getpan() {
    try {
      return _sharedPreferences!.getString('pan') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setsignature(String value) {
    return _sharedPreferences!.setString('sign', value);
  }

  String getsignature() {
    try {
      return _sharedPreferences!.getString('sign') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  String getphoneno() {
    try {
      return _sharedPreferences!.getString('phone_no') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setemail(String value) {
    return _sharedPreferences!.setString('email', value);
  }

  String getemail() {
    try {
      return _sharedPreferences!.getString('email') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setdigilockerurl(String value) {
    return _sharedPreferences!.setString('digilockerurl', value);
  }

  String getdigilockerurl() {
    try {
      return _sharedPreferences!.getString('digilockerurl') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setselfie(String value) {
    return _sharedPreferences!.setString('selfie', value);
  }

  String getselfie() {
    try {
      return _sharedPreferences!.getString('selfie') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setCameraPageDestination(String value) {
    return _sharedPreferences!.setString('destination', value);
  }

  String getCameraPageDestination() {
    try {
      return _sharedPreferences!.getString('destination') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setinappurl(String value) {
    return _sharedPreferences!.setString('inappurl', value);
  }

  String getinappurl() {
    try {
      return _sharedPreferences!.getString('inappurl') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setESignUrl(String value) {
    return _sharedPreferences!.setString('esignurl', value);
  }

  String getESignUrl() {
    try {
      return _sharedPreferences!.getString('esignurl') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setdeeplink(String value) {
    List splitvalue = value.split("/");

    return _sharedPreferences!.setString('deeplink', splitvalue[3]);
  }

  String getdeeplink() {
    try {
      return _sharedPreferences!.getString('deeplink') ?? '';
    } catch (e) {
      return 'initialRoute';
    }
  }

  Future<void> setBiometricAuthStatus(bool value) {
    return _sharedPreferences!.setBool('biometricAuthStatus', value);
  }

  bool getBiometricAuthStatus() {
    try {
      return _sharedPreferences!.getBool('biometricAuthStatus') ?? false;
    } catch (e) {
      return false;
    }
  }

  Future<void> setPinstatus(bool value) {
    return _sharedPreferences!.setBool('pin_status', value);
  }

  bool getPinStatus() {
    try {
      return _sharedPreferences!.getBool('pin_status') ?? false;
    } catch (e) {
      return false;
    }
  }

  Future<void> setauthtoken(String value) {
    return _sharedPreferences!.setString('auth_token', value);
  }

  String getauthtoken() {
    try {
      return _sharedPreferences!.getString('auth_token') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setrefreshtoken(String value) {
    return _sharedPreferences!.setString('refresh_token', value);
  }

  String getrefreshtoken() {
    try {
      return _sharedPreferences!.getString('refresh_token') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setfirstname(String value) {
    return _sharedPreferences!.setString('first_name', value);
  }

  String getfirstname() {
    try {
      return _sharedPreferences!.getString('first_name') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setotpmessage(String value) {
    return _sharedPreferences!.setString('otp_message', value);
  }

  String getotpmessage() {
    try {
      return _sharedPreferences!.getString('otp_message') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setjourney(String value) {
    return _sharedPreferences!.setString('journey', value);
  }

  String getjourney() {
    try {
      return _sharedPreferences!.getString('journey') ?? '';
    } catch (e) {
      return 'firstlogin';
    }
  }

  Future<void> setonboardingAccountNo(String value) {
    return _sharedPreferences!.setString('accountno', value);
  }

  String getAccountNo() {
    try {
      return _sharedPreferences!.getString('accountno') ?? '';
    } catch (e) {
      return 'error';
    }
  }

  Future<void> setonboardingstatus(String value) {
    return _sharedPreferences!.setString('onboardingstatus', value);
  }

  String getonboardingstatus() {
    try {
      return _sharedPreferences!.getString('onboardingstatus') ?? '';
    } catch (e) {
      return 'firstlogin';
    }
  }

  Future<void> setloggedin(bool value) {
    return _sharedPreferences!.setBool('login_status', value);
  }

  bool getloggedin() {
    try {
      return _sharedPreferences!.getBool('login_status') ?? false;
    } catch (e) {
      return false;
    }
  }
}
