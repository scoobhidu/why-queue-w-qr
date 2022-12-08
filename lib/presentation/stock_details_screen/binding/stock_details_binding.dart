import '../controller/stock_details_controller.dart';
import 'package:get/get.dart';

class StockDetailsBinding extends Bindings {
  @override
  void dependencies() {
    Get.lazyPut(() => StockDetailsController());
  }
}
