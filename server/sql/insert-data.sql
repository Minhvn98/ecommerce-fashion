--  --
--
--  XÓA HẾT DỮ LIỆU TRONG CÁC BẢNG NÀY NHÉ.
--
--    INSERT CATEGORIES
--
INSERT INTO
  `categories`(id, name)
VALUES(1, 'Quần'),
  (2, 'Áo'),
  (3, 'Váy'),
  (4, 'Đồ ngủ');
--
  --    INSERT PRODUCTS
  --
  -- Insert products has category is Đồ ngủ (id = 4)
  --
INSERT INTO
  `products`(
    id,
    name,
    description,
    price,
    sale_percent,
    category_id,
    quantity
  )
VALUES(
    1,
    'SET ĐỒ NGỦ ULZZANG BỒ CÔNG ANH VẢI MÁT',
    'Ui khách lưu ý hiện có nhiều shop bán hàng đểu mà lấy hình mẫu giống bên mình, nhưng chất lượng rất tồi. Hãy tỉnh táo nhé. Hãy đọc feedback của khách bên họ để xem vải đểu thế nào nha, hoặc qua bên Feed của shop mình để đọc bài so sánh nhé. Hãy là người mua hàng thông minh nha khách iu!!!',
    169000,
    0,
    4,
    100
  ),(
    2,
    'Đồ bộ mặc nhà dễ thương JANE họa tiết dâu HD04',
    'Đường may chắc chắn và tỉ mỉ đảm bảo trải nghiệm tốt nhất khi sử dụng sản phẩm.
    Công dụng: đem lại sự xinh đẹp và thoải mái nhất cho khách hàng.',
    169000,
    0,
    4,
    120
  ),(
    3,
    'Mùa hè địu nữ váy ngủ mùa hè mỏng dệt bông dễ thương ngọt ngào nơ phong cách cổ tích váy ngủ không tay',
    'Đường may chắc chắn và tỉ mỉ đảm bảo trải nghiệm tốt nhất khi sử dụng sản phẩm.
    Công dụng: đem lại sự xinh đẹp và thoải mái nhất cho khách hàng.',
    269000,
    0,
    4,
    110
  ),(
    4,
    'Váy ngủ 2 dây đầm ngủ thun mềm mịn, thoáng mát [VN01]',
    'Chất liệu: lụa Pháp cao cấp, mềm mại, thoáng mát, mỏng mịn, êm dịu với làn da, các mẹ bỉm sữa có thể ôm con thoải mái khi mặc sản phẩm này mà không lo sẽ làm làn da bé bị kích ứng.',
    149000,
    0,
    4,
    90
  ),(
    5,
    'Đồ ngủ 2 dây nữ có đệm ngực bộ đồ mặc nhà mặc đi biển dễ thương [BC18]',
    'Chất liệu: lụa Pháp cao cấp, mềm mại, thoáng mát, mỏng mịn, êm dịu với làn da, các mẹ bỉm sữa có thể ôm con thoải mái khi mặc sản phẩm này mà không lo sẽ làm làn da bé bị kích ứng.',
    179000,
    0,
    4,
    99
  ),(
    6,
    'Bộ đồ mặc nhà chất liệu đũi JANE mã D101',
    ' Thiết kế độc quyền xinh xắn, độc đáo với form pijama cổ điển đơn giản, thoải mái, quần cộc cùng áo cộc cực trẻ trung, mát mẻ cho mùa hè. Họa tiết kẻ xinh xắn, trendy phong cách Hàn Quốc cho các cô nàng thỏa sức lựa chọn.',
    189000,
    0,
    4,
    69
  ),(
    7,
    'Đồ ngủ - Set Pijama lụa JANE quần cộc, áo cộc họa tiết dễ thương phong cách Hàn Quốc LP10',
    'Chất liệu: lụa Pháp cao cấp, mềm mại, thoáng mát, mỏng mịn, êm dịu với làn da, các mẹ bỉm sữa có thể ôm con thoải mái khi mặc sản phẩm này mà không lo sẽ làm làn da bé bị kích ứng.
    Đường may chắc chắn và tỉ mỉ đảm bảo trải nghiệm tốt nhất khi sử dụng sản phẩm.',
    299000,
    0,
    4,
    96
  ),(
    8,
    'Bộ đồ ngủ lụa Pháp JANE nơ cổ mã LP40',
    ' Thiết kế độc quyền xinh xắn, độc đáo với form pijama cổ điển đơn giản, thoải mái, quần cộc cùng áo cộc cực trẻ trung, mát mẻ cho mùa hè. Họa tiết kẻ xinh xắn, trendy phong cách Hàn Quốc cho các cô nàng thỏa sức lựa chọn.',
    269000,
    0,
    4,
    96
  ),(
    9,
    'Bộ đồ thiết kế cổ vuông viền bèo áo dài quần dài JANE chất liệu vài đũi',
    'Chất liệu: lụa Pháp cao cấp, mềm mại, thoáng mát, mỏng mịn, êm dịu với làn da, các mẹ bỉm sữa có thể ôm con thoải mái khi mặc sản phẩm này mà không lo sẽ làm làn da bé bị kích ứng.
    Đường may chắc chắn và tỉ mỉ đảm bảo trải nghiệm tốt nhất khi sử dụng sản phẩm.',
    119000,
    0,
    4,
    20
  ),(
    10,
    'Set đồ mặc nhà vải mềm dễ thương ( có ảnh thật )',
    'Chất liệu: lụa Pháp cao cấp, mềm mại, thoáng mát, mỏng mịn, êm dịu với làn da, các mẹ bỉm sữa có thể ôm con thoải mái khi mặc sản phẩm này mà không lo sẽ làm làn da bé bị kích ứng.
     Thiết kế độc quyền xinh xắn, độc đáo với form pijama cổ điển đơn giản, thoải mái, quần cộc cùng áo cộc cực trẻ trung, mát mẻ cho mùa hè. Họa tiết kẻ xinh xắn, trendy phong cách Hàn Quốc cho các cô nàng thỏa sức lựa chọn.',
    129000,
    0,
    4,
    10
  );
--
  --
  --
  --
  -- Insert products has category is Áo (id = 2)
  --
  --
  --
INSERT INTO
  `products`(
    id,
    name,
    description,
    price,
    sale_percent,
    category_id,
    quantity
  )
VALUES(
    11,
    'Áo croptop dệt kim tay ngắn phối họa tiết sọc ngang trẻ trung cho nữ',
    'Chất liệu cao cấp, mềm mại, thoáng khí, chống vón cục, sờ vào rất mịn.
    Thiết kế độc đáo giúp bạn trở nên thanh lịch và thời trang mỗi ngày.',
    169000,
    0,
    2,
    100
  ),(
    12,
    'Áo sơ mi nữ cổ vest cộc tay hai túi ngực kiểu nắp, Áo sơ mi nữ cổ vest tay ngắn trơn thiết kế 2 túi nắp ngực Hàn Quốc',
    'Chất liệu cao cấp, mềm mại, thoáng khí, chống vón cục, sờ vào rất mịn.
    Thiết kế độc đáo giúp bạn trở nên thanh lịch và thời trang mỗi ngày.',
    159000,
    0,
    2,
    130
  ),(
    13,
    'Áo Croptop Dệt Kim Tay Ngắn Cổ Tròn Phong Cách Hàn Quốc Cho Nữ',
    'Tụi mình vừa ra mắt mẫu áo rất là xinh và hợp với thời tiết mùa này lắm đây khách ơiii , Áo form siêu tôn dáng và dễ mặc lại giữ ấm tốt lắm ý (À các bạn ở Sài Gòn cũng đừng lo bị nóng nha mặc dù là tay dài nhưng áo có độ thấm hút mồ hôi mặc thoải mái lắm nè ^^).',
    159000,
    0,
    2,
    55
  ),(
    14,
    'Áo Croptop Kiểu Nữ - Áo Crt Khuy Cài Siêu Hot - LP.STORE',
    'Tụi mình vừa ra mắt mẫu áo rất là xinh và hợp với thời tiết mùa này lắm đây khách ơiii , Áo form siêu tôn dáng và dễ mặc lại giữ ấm tốt lắm ý (À các bạn ở Sài Gòn cũng đừng lo bị nóng nha mặc dù là tay dài nhưng áo có độ thấm hút mồ hôi mặc thoải mái lắm nè ^^).',
    159000,
    0,
    2,
    55
  ),(
    15,
    'Áo croptop nữ trễ vai tay ngắn kiểu ôm, áo crt ôm đắp chéo ngực siêu dễ thương',
    'Tụi mình vừa ra mắt mẫu áo rất là xinh và hợp với thời tiết mùa này lắm đây khách ơiii , Áo mỏng nhẹ nên mặc rất dễ chịu, thoải mái, ngoài ra còn che được 1 số khuyết điểm cơ thể .',
    259000,
    0,
    2,
    55
  ),(
    16,
    'Áo Khoác Cardigan Len Mỏng Cho Nữ Phong Cách Hàn Quốc Sexy, Labibi',
    'Áo khoác Cardigan với phong cách trẻ trung, thiết kế đơn giản dễ phối đồ, có thể kết hợp với nhiều loại trang phục như: váy 2 dây, váy body, áo croptop, bikini,… Phù hợp cho mọi cô gái, đặc biệt là những cô nàng thích style nhẹ nhàng quyến rũ',
    259000,
    0,
    2,
    15
  ),(
    17,
    'Áo Khoác Cardigan Len Mỏng Phong Cách Hàn Quốc AK01',
    'Áo khoác Cardigan với phong cách trẻ trung, thiết kế đơn giản dễ phối đồ, có thể kết hợp với nhiều loại trang phục như: váy 2 dây, váy body, áo croptop, bikini,… Phù hợp cho mọi cô gái, đặc biệt là những cô nàng thích style nhẹ nhàng quyến rũ',
    229000,
    0,
    2,
    44
  ),(
    18,
    'Áo croptop tay dài cột nơ màu trắng siêu xinh, siêu đẹp',
    'Áo thiết kế dáng dáng ôm quyến rũ,  hai dây có chốt điều chỉnh để có thể điều chỉnh độ dài của áo và độ trễ của ngực. Áo có thể kế hợp với quần jean, chân váy mang đến sự trẻ trung, năng động hay kết hợp với áo vest, áo khoác thu đông cũng rất hợp',
    229000,
    0,
    2,
    44
  ),(
    19,
    'Áo hai dây Mandela croptop nữ dáng ôm siêu xịn xò, áo 2 dây nữ croptop đẹp vải cotton co dãn',
    'Áo hai dây không thể thiếu trong tủ đồ của các nàng, vừa mát mẻ lại cứ diện vào là “xinh hết nấc”. Áo thiết kế dáng dáng ôm quyến rũ,  hai dây có chốt điều chỉnh để có thể điều chỉnh độ dài của áo và độ trễ của ngực.',
    189000,
    0,
    2,
    64
  ),(
    20,
    'ÁO 2 DÂY BASIC VINTAGE SIÊU HOT 2020 ( kèm ảnh thật )',
    'Áo có thể kế hợp với quần jean, chân váy mang đến sự trẻ trung, năng động hay kết hợp với áo vest, áo khoác thu đông cũng rất hợp. Chắc chắn bạn không thể bỏ qua chiếc áo 2 dây đẹp này rồi!',
    189000,
    0,
    2,
    72
  );
--
  --
  --
  --
  -- Insert products has category is Váy (id = 3)
  --
  --
  --
INSERT INTO
  `products`(
    id,
    name,
    description,
    price,
    sale_percent,
    category_id,
    quantity
  )
VALUES(
    21,
    'Váy công chúa trễ vai ren ulzzang - V15 - jannahouse',
    'Siêu phẩm váy cột dây cho các tiểu thư bánh bèo mới cập kho nhà AnieStore nha!! Váy suông hồng họa tiết hoa nhí.
    Các nàng xuống phố, dạo biển, vi vu du lịch cực hợp đó nha!',
    169000,
    0,
    3,
    20
  ),(
    22,
    'Váy hoa ullzang dáng dài cổ tim voan tơ - Đầm hoa nhí Vintage',
    'Các nàng xuống phố, dạo biển, vi vu du lịch cực hợp đó nha!
    Phối nhẹ vs chiếc túi vải Canvas nhà anie nữa thì hết í đó ạ',
    139000,
    0,
    3,
    22
  ),(
    23,
    'Đầm nữ cổ V họa tiết dập nổi thanh lịch ngọt ngào nữ tính',
    'Các nàng xuống phố, dạo biển, vi vu du lịch cực hợp đó nha!
    Phối nhẹ vs chiếc túi vải Canvas nhà anie nữa thì hết í đó ạ',
    149000,
    0,
    3,
    33
  ),(
    24,
    'Đầm - váy hai dây voan tơ buộc vai 2 lớp ',
    'Mẫu váy nhẹ nhàng tiểu thư cho các nàng bánh bèo vừa về kho Vogue nha!
    Mã mới xưởng bên anie mới thiết kế chào hàng các nàng luôn ạ
    Thiết kế nơ ngưc buộc, tay bồng dáng dài xinh lắm nhé',
    179000,
    0,
    3,
    35
  ),(
    25,
    'Đầm Nữ Cổ Vuông Phối Nút siêu xinh nhé các nàng',
    'Thiết kế nơ ngưc buộc, tay bồng dáng dài xinh lắm nhé. Vải đũi xốp trắng mịn dày dặn.
    Dáng váy rất dễ phối với các phụ kiện nha, giày boots vs guốc cực hợp đó ạ',
    189000,
    0,
    3,
    45
  ),(
    26,
    'Váy Hoa Hai Dây Dáng Dài Kiểu Dáng Dễ Thương Xinh Xắn',
    'Thiết kế nơ ngưc buộc, tay bồng dáng dài xinh lắm nhé. Vải đũi xốp trắng mịn dày dặn.
    Dáng váy rất dễ phối với các phụ kiện nha, giày boots vs guốc cực hợp đó ạ',
    189000,
    0,
    3,
    5
  ),(
    27,
    'Đầm váy trắng cổ bèo đan dây trước (kèm hình thật)',
    'Mẫu váy nhẹ nhàng tiểu thư cho các nàng bánh bèo vừa về kho Vogue nha!
    Mã mới xưởng bên anie mới thiết kế chào hàng các nàng luôn ạ.
    Thiết kế nơ ngưc buộc, tay bồng dáng dài xinh lắm nhé.
    Vải đũi xốp trắng mịn dày dặn',
    239000,
    0,
    3,
    11
  ),(
    28,
    'Đầm Nữ Babydoll Nơ Ngực xinh xắn tiểu thư cực kỳ duyên dáng khi đi học và đi chơi Mã VT02',
    'Mẫu váy nhẹ nhàng tiểu thư cho các nàng bánh bèo vừa về kho Vogue nha!
    Mã mới xưởng bên anie mới thiết kế chào hàng các nàng luôn ạ.
    Thiết kế nơ ngưc buộc, tay bồng dáng dài xinh lắm nhé.
    Vải đũi xốp trắng mịn dày dặn',
    239000,
    0,
    3,
    11
  ),(
    29,
    'Váy hai dây hoa nhí - Đầm bông hai dây Maina',
    'Váy nhà Janna mặc được 4 mùa luôn các nàng nhé ,,
    Thích hợp mặc đi học, đi dạo phố, chụp ảnh sống ảo party bất chấp luôn ,,,
    Đảm bảo mặc vào xinh hơn đứa mình ghét ạ >"<',
    259000,
    0,
    3,
    33
  ),(
    30,
    'Váy ulzzang tiểu thư công chúa phối nơ - V16 - jannahouse',
    'Váy nhà Janna mặc được 4 mùa luôn các nàng nhé ,,
    Thích hợp mặc đi học, đi dạo phố, chụp ảnh sống ảo party bất chấp luôn ,,,
    Đảm bảo mặc vào xinh hơn đứa mình ghét ạ >"<',
    269000,
    0,
    3,
    66
  );
--
  --
  --
  --
  -- Insert products has category is Quần (id = 1)
  --
  --
  --
INSERT INTO
  `products`(
    id,
    name,
    description,
    price,
    sale_percent,
    category_id,
    quantity
  )
VALUES(
    31,
    'Quần Sooc Quần Short Nữ Vải Xước X.0 Hàn Chất Vải Thoáng Mát Phong Cách Thời Trang Hàn Quốc',
    'Quần short được yêu chuộng bởi nó tôn lên cặp chân dài miên man của cô gái, rất thoáng mát khi sử dụng và độ thời trang không hề thua kém bất cứ một item nào khác. Với quần short, những ngày đầy nắng của nàng sẽ thêm thú vị và tràn ngập niềm vui.',
    169000,
    0,
    1,
    20
  ),(
    32,
    'Quần Short Cạp Cao Ulzzang - Hàng có sẵn',
    'Chị em muốn tìm loại quần sooc mềm , kín đáo , không quá hở hay sexy mà mặc vẫn thoải mái + tiện lợi + mix với áo nào cũng đước thì đến ngay với e sooc ĐŨI XƯỚC này của e nhé ',
    169000,
    0,
    1,
    33
  ),(
    33,
    'Quần short ống rộng CRQ023 Unisex Nam - Nữ cạp cao lưng thun màu đen D50',
    'Chị em muốn tìm loại quần sooc mềm , kín đáo , không quá hở hay sexy mà mặc vẫn thoải mái + tiện lợi + mix với áo nào cũng đước thì đến ngay với e sooc ĐŨI XƯỚC này của e nhé ',
    169000,
    0,
    1,
    53
  ),(
    34,
    'Quần SHORT JEAN 01 Nữ 1hitshop (màu XANH NHẠT)',
    'Giới thiệu sản phẩm Quần ống rộng ĐŨI XƯỚC cạp chun siêu đẹp
    Chiếc quần dài ống rộng đũi này sẽ khiến tủ đồ của bạn thêm nhiều lựa chọn dạo phố nha. Sản phẩm cực kỳ thoáng mát, giúp bạn năng động suốt ngày',
    179000,
    0,
    1,
    69
  ),(
    35,
    'Quần giả váy phối cúc siêu Hot - HKA001',
    'Quần giả váy cực xinh luôn ạ, Form bao đẹp khách nhé.Hàng nhà e loại 1 ạ
    Lên dáng rất ok',
    179000,
    0,
    1,
    69
  ),(
    36,
    'Quần Đũi Dài quần ống rộng siêu Hot cao cấp mẫu mới 2021',
    'Chất liệu đũi linen (thoáng mát, nhẹ và thấm hút mồ hôi cực kỳ tốt , sợi vải bền , chắc , càng dùng lâu vải càng mềm và êm.)',
    179000,
    0,
    1,
    59
  ),(
    37,
    'Quần Đũi Dài quần ống rộng siêu Hot cao cấp mẫu mới 2021',
    'Chất liệu đũi linen (thoáng mát, nhẹ và thấm hút mồ hôi cực kỳ tốt , sợi vải bền , chắc , càng dùng lâu vải càng mềm và êm.)',
    179000,
    0,
    1,
    69
  ),(
    38,
    'Quần Ống Rộng Nữ vải Đũi cạp chun MADELA siêu hack dáng kute, Quần Ống Suông Nữ chất Đũi mát',
    'Chất liệu đũi linen (thoáng mát, nhẹ và thấm hút mồ hôi cực kỳ tốt , sợi vải bền , chắc , càng dùng lâu vải càng mềm và êm.)',
    179000,
    0,
    1,
    69
  ),(
    39,
    'Quần đũi linen 2 màu siêu xinh, siêu đẹp nhé',
    'Đặc điểm: vải rất bền và đẹp cùng kiểu dáng culottes thời thượng giúp chị em tha hồ kết hợp style thời trang mới lạ cho riêng mình, ',
    229000,
    0,
    1,
    69
  ),(
    40,
    'Quần Đùi In Gấu Nâu Unisex Hot Hit Quần Đùi Ống Rộng Trendy',
    'Đặc điểm: vải rất bền và đẹp cùng kiểu dáng culottes thời thượng giúp chị em tha hồ kết hợp style thời trang mới lạ cho riêng mình, ',
    249000,
    0,
    1,
    69
  );
