// import 'package:mvp_draft_4/core/app_export.dart';
// import 'package:mvp_draft_4/data/apiClient/api_client.dart';
//
// import 'helper_utils.dart';
//
// class InitialBindings extends Bindings {
//   @override
//   Future<void> dependencies() async {
//     Get.put(PrefUtils());
//     Get.put(ApiClient());
//     Get.put(HelperUtils());
//     Get.put<bool>(false, tag: 'aadharNotLinkedWithPan');
//
//     bool status = await Get.find<HelperUtils>().request_CameraPermission();
//     // Get.put<bool>(false, tag: 'aadharNotLinkedWithPan');
//     Connectivity connectivity = Connectivity();
//     Get.put(NetworkInfo(connectivity));
//   }
// }
