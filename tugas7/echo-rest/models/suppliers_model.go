package models

import (
	"fmt"
	"net/http"
	"tugas7/echo-rest/common"
	"tugas7/echo-rest/db"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

var errMessage string
var errs *ex.AppError
var cust common.Suppliers
var custObj []common.Suppliers

func FetchSuppliers() (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	con := db.CreateCon()

	sqlQuery := `SELECT
					IFNULL(CompanyName,''), CompanyName
					IFNULL(ContactName,'') ContactName,
					IFNULL(ContactTitle,'') ContactTitle,
					IFNULL(Address,'') Address,
					IFNULL(PostalCode,'') PostalCode
				FROM suppliers `

	rows, err := con.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&cust.CompanyName, &cust.ContactName, &cust.ContactTitle, &cust.Address, &cust.PostalCode)

		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}

		custObj = append(custObj, cust)

	}

	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = custObj

	return res, nil
}

//AddSuppliers ...
func AddSuppliers(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `INSERT INTO suppliers (CompanyName,ContactName,ContactTitle,Address,PostalCode)
					 VALUES (?,?,?,?,?)`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, req.Address, req.PostalCode)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"Suppliers ADD": req.CompanyName,
	}

	return res, nil
}

//UpdateSuppliers ...
func UpdateSuppliers(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `UPDATE suppliers SET CompanyName = ?, ContactName = ?, ContactTitle = ?, Address = ?, PostalCode = ? WHERE  SupplierID = '` + req.CompanyName + `'`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, req.Address, req.PostalCode)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"row affected :": req.CompanyName,
	}

	return res, nil
}

//DeleteSuppliers ...
func DeleteSuppliers(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `DELETE FROM suppliers WHERE  SupplierID = ?`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"row deleted :": req.CompanyName,
	}

	return res, nil
}
