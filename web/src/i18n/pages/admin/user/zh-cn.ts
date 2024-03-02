export default {
  adminUser: {
    userInfo:"用户信息",
    query:"查询",
    add_user:"新增用户",
    modify_user:"修改用户",
    advanced_query:"高级查询",
    customerService:"客户服务",
    SysUser:{
      index:"序号",
      created_at: "创建时间",
      updated_at: "更新时间",
      id:"id",
      uuid: "uuid",
      user_name: "用户名",
      nick_name: "昵称",
      password: "密码",
      avatar: "头像",
      enable: "状态",
      invitation_code: "邀请码",
      referrer_code: "推荐人的邀请码",
      balance: "余额",
      tg_id: "tg id",
      role_group: "角色组",	//角色组
      orders: "订单",     //订单
    },
    CustomerService: {
      index:"序号",
      created_at:"创建时间",
      updated_at:"更新时间",
      id:"id",
      user_id:"用户id",
      user_name:"用户名",
      service_status:"服务状态",
      service_start_at:"服务开始时间",
      service_end_at:"服务结束时间",
      is_renew:"是否可续费",
      renewal_amount:"续费价格",
      goods_id:"商品id",
      subject:"商品标题",
      des:"商品描述",
      price:"商品价格",
      goods_type:"商品类型",
      duration:"订购时长",
      total_bandwidth:"总流量",
      node_connector:"节点连接数",
      node_speed_limit:"节点速率",
      traffic_reset_day:"流量重置日",
      sub_status: "订阅状态",
      sub_uuid:"订阅UUID",
      used_up:"已用上行",
      used_down:"已用下行",
    },
  },
};