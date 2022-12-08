// import 'dart:async';
// import 'package:connectivity_plus/connectivity_plus.dart';
//
// // For checking internet connectivity
// abstract class NetworkInfoI {
//   Future<bool> isConnected();
//
//   Future<ConnectivityResult> get connectivityResult;
//
//   Stream<ConnectivityResult> get onConnectivityChanged;
// }
//
// class NetworkInfo implements NetworkInfoI {
//   Connectivity connectivity;
//
//   late Rx<bool> disableScreen = false.obs;
//   late Connectivity _connectivity;
//   late StreamSubscription<ConnectivityResult> connectivitySubscription;
//   Rx<ConnectivityResult> connectStatus = ConnectivityResult.wifi.obs;
//
//   NetworkInfo(this.connectivity) {
//     connectivity = this.connectivity;
//     _connectivity = this.connectivity;
//     initialConnectionCheck();
//
//     connectivitySubscription = _connectivity.onConnectivityChanged.listen((event) {
//       connectStatus.value = event;
//
//       if (connectStatus.value == ConnectivityResult.mobile ||
//           connectStatus.value == ConnectivityResult.wifi ||
//           connectStatus.value == ConnectivityResult.ethernet) {
//         print('Connnected via ${connectStatus.value}');
//         disableScreen.value = false;
//         print("Working");
//       } else if (connectStatus.value == ConnectivityResult.none ||
//           connectStatus.value == ConnectivityResult.bluetooth) {
//         disableScreen.value = true;
//         print('Not Connnected');
//         print('Screen Disabled');
//       }
//       else{
//         disableScreen.value = true;
//         print('Not Connnected');
//         print('Screen Disabled');
//       }
//     });
//   }
//
//   void initialConnectionCheck() async {
//     connectStatus.value = await _connectivity.checkConnectivity();
//
//     if (connectStatus.value == ConnectivityResult.mobile ||
//         connectStatus.value == ConnectivityResult.wifi ||
//         connectStatus.value == ConnectivityResult.ethernet) {
//       print('Connnected via ${connectStatus.value}');
//       disableScreen.value = false;
//       print("Working");
//     } else if (connectStatus.value == ConnectivityResult.none ||
//         connectStatus.value == ConnectivityResult.bluetooth) {
//       disableScreen.value = true;
//       print('Not Connnected');
//       print('Screen Disabled');
//     } else{
//       disableScreen.value = true;
//       print('Not Connnected');
//       print('Screen Disabled');
//     }
//   }
//   ///checks internet is connected or not
//   ///returns [true] if internet is connected
//   ///else it will return [false]
//   @override
//   Future<bool> isConnected() async {
//     final result = await connectivity.checkConnectivity();
//     if (result != ConnectivityResult.none) {
//       return true;
//     }
//     return false;
//   }
//
//   // to check type of internet connectivity
//   @override
//   Future<ConnectivityResult> get connectivityResult async {
//     return connectivity.checkConnectivity();
//   }
//
//   //check the type on internet connection on changed of internet connection
//   @override
//   Stream<ConnectivityResult> get onConnectivityChanged =>
//       connectivity.onConnectivityChanged;
// }