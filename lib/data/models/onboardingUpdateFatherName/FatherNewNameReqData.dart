// search JSON to dart to create a models for API request or response like this

class FatherNewNameReqData {
  FatherNewNameReqData({
    required this.phone,
    required this.new_name,
  });

  late final int phone;
  late final String new_name;

  FatherNewNameReqData.fromJson(Map<String, dynamic> json){
  phone = json['phone'];
  new_name = json['new name'];
  }

  Map<String, dynamic> toJson() {
    final _data = <String, dynamic>{};
    _data['phone'] = phone;
    _data['new name'] = new_name;
    return _data;
  }
}