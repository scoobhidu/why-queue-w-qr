import 'dart:convert';
import 'dart:developer';

import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;

import 'dart:io';

import '../../core/errors/exceptions.dart';
import '../../core/network/network_info.dart';
import '../../core/utils/image_constant.dart';
import '../../core/utils/pref_utils.dart';
import '../../core/utils/progress_dialog_utils.dart';

class ApiClient extends GetConnect {
  var url = dotenv.env['url'];
  var ssourl = dotenv.env['ssourl'];
  String xapikey = dotenv.env['xapikey']!;
  var cashDummyUrl = dotenv.env['cashDummyUrl'];
  String x_api_keyCash = dotenv.env['x_api_keyCash']!;
  var cashLiveURL = "http://cash-api.test.ventura1.com:8001";

  @override
  void onInit() {
    super.onInit();
    httpClient.timeout = Duration(seconds: 60);

    httpClient.addRequestModifier<dynamic>((request) {
      Map<String, String> headers = {
        "Content-type": "application/json",
      };

      request.headers.addAll(headers);

      return request;
    });
  }

  ///method can be used for checking internet connection
  ///returns [bool] based on availability of internet
  ///

  /// is `true` when the response status code is between 200 and 299
  ///
  /// user can modify this method with custom logics based on their API response
  bool _isSuccessCall(Response response) {
    return response.isOk;
  }

  Future createOtpVerify({Function(dynamic data)? onSuccess,
    Function(dynamic error)? onError,
    Map<String, String> headers = const {},
    Map requestData = const {}}) async {
    String sessionid = await Get.find<PrefUtils>().getsessionid();
    headers = {"session_id": sessionid.toString()};
    ProgressDialogUtils.showProgressDialog();
    print(requestData);
    try {
      // await isNetworkConnected();
      Response response = await httpClient.post(
          '$url/onboarding/v2/signup/user/phone/otp/verify',
          headers: headers,
          body: requestData);
      ProgressDialogUtils.hideProgressDialog();
      print(response.body);
      print(requestData);
      if (_isSuccessCall(response)) {
        onSuccess!(response.body);
      } else {
        print(response.statusCode);
        onError!(
          response.hasError
              ? response.body!["message"]
              : 'Something Went Wrong!',
        );
      }
    } catch (error) {
      ProgressDialogUtils.hideProgressDialog();
      onError!(error);
    }
  }
}
