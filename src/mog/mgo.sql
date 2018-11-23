// 插入 inventory集合
//db.inventory.insertMany([
//   { item: "journal", qty: 25, tags: ["blank", "red"], size: { h: 14, w: 21, uom: "cm" } },
//   { item: "mat", qty: 85, tags: ["gray"], size: { h: 27.9, w: 35.5, uom: "cm" } },
//   { item: "mousepad", qty: 25, tags: ["gel", "blue"], size: { h: 19, w: 22.85, uom: "cm" } }
//])

// 查询所有
//db.inventory.find({})

// select * from inventory where item in("journal","mat")
//db.inventory.find({
//    item: {
//        $in: ["journal", "mat"]
//    }
//})

// select * from inventory where qty < 26;  $gt $lte $gte
//db.inventory.find({
//	qty: { $lt: 26}
//})

// select * find inventory where qty < 30 and item = "journal"
//db.inventory.find({
//	qty: {$lt: 30},
//	item: "journal"
//})

// select * find inventory where status = "A" or qty < 30;
//db.inventory.find({
//    $or: [{
//        status: "A"
//    }, {
//        qty: {
//            $lt: 30
//        }
//    }]
//})

// 查询集合子对象字段
//db.inventory.find({
//		"size.h": 19
//})
