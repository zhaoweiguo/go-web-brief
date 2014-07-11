package main

import (
	"fmt"
	"github.com/zhaoweiguo/go-web-brief/config"
	"github.com/zhaoweiguo/go-web-brief/db"
	"github.com/zhaoweiguo/go-web-brief/error"
	"github.com/zhaoweiguo/go-web-brief/utils"

)

func renderPages(page int) FinalData {

	fmt.Println(page)

	var finalData FinalData

	var mainDataList []MainData
	var pageDataList []string

	mainDataList = getMainData(page)
	pageDataList = getPageData()

	finalData.MainData = mainDataList
	finalData.PageData = pageDataList
	return finalData
}

func getPageData() []string {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
}

func getMainData(page int) []MainData {
	dateData := getDateData(page)
	mainItems := getMainItem(page)
	mainData := MainData{dateData, mainItems}

	return []MainData {	mainData }
}

func getDateData(page int) string{
//	return "2004-10-20 Monday"
	return utils.DisplayBeforeDate(page-1);
}


func getMainItem(page int) []MainItem {
	date := utils.BeforeDate(page-1);
	rows := db.QueryByDate(config.TAB_MAIN, date)

	var id int
	var title string
	var image string
	var url string

	mainData := make([]MainItem, 100)
	i := 0;
	for rows.Next() {
		err := rows.Scan(&id, &title, &image, &url)
		error.CheckErr(err)
//		fmt.Println(id, title, pic, url)
		
		mainData[i] = MainItem{id, title, image, url}
		i++
	}

//	fmt.Println(mainData)

	return mainData
}

func renderPages_test(page int) FinalData {
	mainItems := make([]MainItem, 3)
	mainDatas := make([]MainData, 3)
	pageData := []string{"1", "2"}

	mainItem1 := MainItem{1, "abc", "http://zhihudaily.ahorn.me/img/croped/http-__pic2.zhimg.com_2608dc52e50099943ffa7675b225dd88.jpg",  "http://abc.com"}
	mainItem2 := MainItem{2, "abc2", "http://zhihudaily.ahorn.me/img/croped/http-__pic2.zhimg.com_2608dc52e50099943ffa7675b225dd88.jpg",  "http://abc.com"}
	mainItems[0] = mainItem1
	mainItems[1] = mainItem2

	mainData := MainData{"2014-10-20 Friday", mainItems}
	mainDatas[0] = mainData

	finalData := FinalData{mainDatas, pageData}
	return finalData
}

