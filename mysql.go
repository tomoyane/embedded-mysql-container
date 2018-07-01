package main

type Mysql struct {
	driverName string
	dataSourceName string
}

func (m Mysql) SetDriverName(driverName string) {
	m.driverName = driverName
}

func (m Mysql) GetDriverName() string {
	return driverName
}

func (m Mysql) SetDataSourceName(dataSourceName string) {
	m.dataSourceName = dataSourceName
}

func (m Mysql) GetDataSourceName() string {
	return m.dataSourceName
}
