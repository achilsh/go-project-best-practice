* 独立模块: 订单业务的核心逻辑部分, eg: 订单商品服务私有包; 不同类型的业务逻辑代码.  
*  本模块（本子目录）内代码可以相互调用，跨模块需要调用公共的code. 
* eg:存在公共部分则放在pkg部分.