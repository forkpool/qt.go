package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qsortfilterproxymodel.h
// dst-file: /src/core/qsortfilterproxymodel.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QSortFilterProxyModel)=1
type QSortFilterProxyModel struct {
  /*qbase*/ QAbstractProxyModel;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSortFilterProxyModel) setFilterRegExp(args ...interface{}) () {
  // setFilterRegExp(const class QString &)
  // setFilterRegExp(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel15setFilterRegExpERK7QString
  case 1:
    // invoke: _ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRegExp", args)
 }

}


func (this *QSortFilterProxyModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "rowCount", args)
 }

}


func (this *QSortFilterProxyModel) sibling(args ...interface{}) () {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sibling", args)
 }

}


func (this *QSortFilterProxyModel) span(args ...interface{}) () {
  // span(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel4spanERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "span", args)
 }

}


func (this *QSortFilterProxyModel) mapFromSource(args ...interface{}) () {
  // mapFromSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapFromSource", args)
 }

}


func (this *QSortFilterProxyModel) setFilterWildcard(args ...interface{}) () {
  // setFilterWildcard(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel17setFilterWildcardERK7QString
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterWildcard", args)
 }

}


func (this *QSortFilterProxyModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "hasChildren", args)
 }

}


func (this *QSortFilterProxyModel) setFilterFixedString(args ...interface{}) () {
  // setFilterFixedString(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterFixedString", args)
 }

}


func (this *QSortFilterProxyModel) setData(args ...interface{}) () {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setData", args)
 }

}


func (this *QSortFilterProxyModel) setSortRole(args ...interface{}) () {
  // setSortRole(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel11setSortRoleEi
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortRole", args)
 }

}


func (this *QSortFilterProxyModel) data(args ...interface{}) () {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel4dataERK11QModelIndexi
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "data", args)
 }

}


func (this *QSortFilterProxyModel) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel10invalidateEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "invalidate", args)
 }

}


func (this *QSortFilterProxyModel) sortColumn(args ...interface{}) () {
  // sortColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel10sortColumnEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortColumn", args)
 }

}


func (this *QSortFilterProxyModel) insertRows(args ...interface{}) () {
  // insertRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertRows", args)
 }

}


func (this *QSortFilterProxyModel) filterKeyColumn(args ...interface{}) () {
  // filterKeyColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel15filterKeyColumnEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterKeyColumn", args)
 }

}


func (this *QSortFilterProxyModel) canFetchMore(args ...interface{}) () {
  // canFetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "canFetchMore", args)
 }

}


func (this *QSortFilterProxyModel) isSortLocaleAware(args ...interface{}) () {
  // isSortLocaleAware()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel17isSortLocaleAwareEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "isSortLocaleAware", args)
 }

}


func (this *QSortFilterProxyModel) fetchMore(args ...interface{}) () {
  // fetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "fetchMore", args)
 }

}


func (this *QSortFilterProxyModel) mapSelectionFromSource(args ...interface{}) () {
  // mapSelectionFromSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionFromSource", args)
 }

}


func (this *QSortFilterProxyModel) mapSelectionToSource(args ...interface{}) () {
  // mapSelectionToSource(const class QItemSelection &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelection{}) // "const QItemSelection &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapSelectionToSource", args)
 }

}


func (this *QSortFilterProxyModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel9mimeTypesEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mimeTypes", args)
 }

}


func (this *QSortFilterProxyModel) buddy(args ...interface{}) () {
  // buddy(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel5buddyERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "buddy", args)
 }

}


func (this *QSortFilterProxyModel) filterRole(args ...interface{}) () {
  // filterRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel10filterRoleEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRole", args)
 }

}


func (this *QSortFilterProxyModel) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel5clearEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "clear", args)
 }

}


func (this *QSortFilterProxyModel) setFilterKeyColumn(args ...interface{}) () {
  // setFilterKeyColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel18setFilterKeyColumnEi
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterKeyColumn", args)
 }

}


func (this *QSortFilterProxyModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "metaObject", args)
 }

}


func (this *QSortFilterProxyModel) sortRole(args ...interface{}) () {
  // sortRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel8sortRoleEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "sortRole", args)
 }

}


func (this *QSortFilterProxyModel) setSortLocaleAware(args ...interface{}) () {
  // setSortLocaleAware(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel18setSortLocaleAwareEb
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSortLocaleAware", args)
 }

}


func (this *QSortFilterProxyModel) mapToSource(args ...interface{}) () {
  // mapToSource(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "mapToSource", args)
 }

}


func (this *QSortFilterProxyModel) setSourceModel(args ...interface{}) () {
  // setSourceModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setSourceModel", args)
 }

}


func NewQSortFilterProxyModel(args ...interface{})() {
}


func (this *QSortFilterProxyModel) removeColumns(args ...interface{}) () {
  // removeColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeColumns", args)
 }

}


func (this *QSortFilterProxyModel) FreeQSortFilterProxyModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "~QSortFilterProxyModel", args)
 }

}


func (this *QSortFilterProxyModel) dynamicSortFilter(args ...interface{}) () {
  // dynamicSortFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel17dynamicSortFilterEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "dynamicSortFilter", args)
 }

}


func (this *QSortFilterProxyModel) insertColumns(args ...interface{}) () {
  // insertColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "insertColumns", args)
 }

}


func (this *QSortFilterProxyModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "columnCount", args)
 }

}


func (this *QSortFilterProxyModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel6parentERK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "parent", args)
 }

}


func (this *QSortFilterProxyModel) filterRegExp(args ...interface{}) () {
  // filterRegExp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel12filterRegExpEv
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "filterRegExp", args)
 }

}


func (this *QSortFilterProxyModel) setFilterRole(args ...interface{}) () {
  // setFilterRole(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel13setFilterRoleEi
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setFilterRole", args)
 }

}


func (this *QSortFilterProxyModel) removeRows(args ...interface{}) () {
  // removeRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "removeRows", args)
 }

}


func (this *QSortFilterProxyModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "index", args)
 }

}


func (this *QSortFilterProxyModel) setDynamicSortFilter(args ...interface{}) () {
  // setDynamicSortFilter(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN21QSortFilterProxyModel20setDynamicSortFilterEb
  default:
    qtrt.ErrorResolve("QSortFilterProxyModel", "setDynamicSortFilter", args)
 }

}

// <= body block end

