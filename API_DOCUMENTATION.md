# Hulio User Service API 接口文档

## 概述

Hulio User Service 是一个用户管理服务，提供用户注册、登录、信息管理等功能。服务运行在 `:8080` 端口，使用 JWT 进行身份认证。

## 基础信息

- **服务地址**: `https://www.hulio88.xyz/user-service`
- **API版本**: `v1`
- **认证方式**: JWT Token (Bearer Token)
- **数据格式**: JSON
- **字符编码**: UTF-8

## 跨域配置

服务已配置跨域支持，允许以下域名访问：

### 本地开发环境
- `http://localhost:3000` - Next.js 默认端口
- `http://localhost:5173` - Vite 默认端口  
- `http://localhost:8080` - 其他常用端口
- `http://localhost:3001` - 其他常用端口
- `http://localhost:4200` - Angular 默认端口

### Vercel 托管域名
- `https://hulio-user-service.vercel.app`
- `https://hulio-user-service-git-main.vercel.app`
- `https://hulio-user-service-git-dev.vercel.app`
- `https://hulio-user-service-git-feature.vercel.app`

### 生产环境域名
- `https://www.hulio88.xyz`
- `https://hulio88.xyz`

## 通用响应格式

### 成功响应
```json
{
  "success": true,
  "code": 200,
  "message": "操作成功",
  "data": {}
}
```

### 错误响应
```json
{
  "success": false,
  "code": 400,
  "message": "错误描述",
  "data": null
}
```

## 状态码说明

| 状态码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权/认证失败 |
| 404 | 资源不存在 |
| 405 | 方法不允许 |
| 500 | 服务器内部错误 |

## 用户状态枚举

| 值 | 状态 | 说明 |
|----|------|------|
| 1 | 正常 | 用户账号正常 |
| 2 | 非活跃 | 用户账号非活跃 |
| 3 | 暂停 | 用户账号被暂停 |
| 4 | 已删除 | 用户账号已删除 |
| 5 | 待审核 | 用户账号待审核 |

## 用户来源枚举

| 值 | 来源 | 说明 |
|----|------|------|
| 1 | Hulio Site | Hulio网站 |

## 接口列表

### 1. 用户注册

**接口地址**: `POST /api/v1/user/register`

**接口描述**: 创建新用户账号

**请求头**:
```
Content-Type: application/json
```

**请求参数**:
```json
{
  "username": "string",     // 用户名（必填）
  "password": "string",     // 密码（必填）
  "nickname": "string",     // 昵称（可选）
  "avatar": "string",       // 头像URL（可选）
  "email": "string",        // 邮箱（可选）
  "phone": "string",        // 手机号（可选）
  "status": 1,              // 用户状态（可选，默认1）
  "from": 1,                // 用户来源（可选，默认0）
  "extinfo": "string"       // 扩展信息（可选）
}
```

**响应示例**:
```json
{
  "success": true,
  "code": 200,
  "message": "操作成功",
  "data": null
}
```

**错误情况**:
- 用户名已存在: `code: 400, message: "用户名已存在"`

---

### 2. 用户登录

**接口地址**: `POST /api/v1/user/login`

**接口描述**: 用户登录获取访问令牌

**请求头**:
```
Content-Type: application/json
```

**请求参数**:
```json
{
  "username": "string",     // 用户名（必填）
  "password": "string"      // 密码（必填）
}
```

**响应示例**:
```json
{
  "success": true,
  "code": 200,
  "message": "操作成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "testuser",
      "nickname": "测试用户",
      "avatar": "",
      "email": "",
      "phone": "",
      "status": 1,
      "from": 1,
      "extinfo": "",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

**错误情况**:
- 用户不存在: `code: 400, message: "用户不存在"`
- 密码错误: `code: 400, message: "密码错误"`

---

### 3. 获取用户信息

**接口地址**: `GET /api/v1/user`

**接口描述**: 获取当前登录用户的详细信息

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**: 无

**响应示例**:
```json
{
  "success": true,
  "code": 200,
  "message": "操作成功",
  "data": {
    "token": "",
    "user": {
      "id": 1,
      "username": "testuser",
      "nickname": "测试用户",
      "avatar": "",
      "email": "",
      "phone": "",
      "status": 1,
      "from": 1,
      "extinfo": "",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

**错误情况**:
- 未授权: `code: 401, message: "未授权"`
- 用户不存在: `code: 400, message: "用户不存在"`

---

### 4. 更新用户信息

**接口地址**: `PUT /api/v1/user`

**接口描述**: 更新当前登录用户的信息（支持部分更新）

**请求头**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**请求参数**:
```json
{
  "nickname": "string",     // 昵称（可选）
  "avatar": "string",       // 头像URL（可选）
  "email": "string",        // 邮箱（可选）
  "phone": "string",        // 手机号（可选）
  "status": 1,              // 用户状态（可选）
  "extinfo": "string"       // 扩展信息（可选）
}
```

**说明**: 只更新非空字段，空字段将被忽略

**响应示例**:
```json
{
  "success": true,
  "code": 200,
  "message": "操作成功",
  "data": {
    "token": "",
    "user": {
      "id": 1,
      "username": "testuser",
      "nickname": "新昵称",
      "avatar": "新头像URL",
      "email": "new@example.com",
      "phone": "13800138000",
      "status": 1,
      "from": 1,
      "extinfo": "新扩展信息",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

**错误情况**:
- 未授权: `code: 401, message: "未授权"`
- 用户不存在: `code: 400, message: "用户不存在"`

---

### 5. 删除用户

**接口地址**: `DELETE /api/v1/user`

**接口描述**: 删除当前登录用户的账号

**请求头**:
```
Authorization: Bearer {token}
```

**请求参数**: 无

**响应示例**:
```json
{
  "success": true,
  "code": 200,
  "message": "操作成功",
  "data": null
}
```

**错误情况**:
- 未授权: `code: 401, message: "未授权"`
- 用户不存在: `code: 400, message: "用户不存在"`

---

### 6. 健康检查

**接口地址**: `GET /health`

**接口描述**: 服务健康状态检查

**请求头**: 无

**请求参数**: 无

**响应示例**:
```json
{
  "status": "ok",
  "service": "hulio-user-service",
  "timestamp": 1704067200
}
```

---

## 认证说明

### JWT Token 格式
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Token 获取
通过登录接口获取 token，后续需要认证的接口都需要在请求头中携带此 token。

### Token 有效期
Token 具有时效性，过期后需要重新登录获取。

---

## 白名单接口

以下接口无需认证即可访问：
- `POST /api/v1/user/login` - 用户登录
- `POST /api/v1/user/register` - 用户注册
- `GET /health` - 健康检查
- `GET /api/v1/mock/panic` - 测试接口

---

## 错误处理

### 常见错误码

| 错误码 | 说明 | 解决方案 |
|--------|------|----------|
| 001 | 用户名已存在 | 更换用户名重新注册 |
| 002 | 用户不存在 | 检查用户名是否正确 |
| 003 | 密码错误 | 检查密码是否正确 |
| 004 | 未授权 | 检查token是否有效 |
| 500 | 服务器内部错误 | 联系技术支持 |

### 错误响应格式
```json
{
  "success": false,
  "code": 001,
  "message": "用户名已存在",
  "data": null
}
```

---

## 使用示例

### cURL 示例

#### 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "nickname": "测试用户",
    "email": "test@example.com"
  }'
```

#### 用户登录
```bash
curl -X POST http://localhost:8080/api/v1/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

#### 获取用户信息
```bash
curl -X GET http://localhost:8080/api/v1/user \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 更新用户信息
```bash
curl -X PUT http://localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "nickname": "新昵称",
    "email": "new@example.com"
  }'
```

---

## 注意事项

1. **密码安全**: 密码在传输前会进行哈希处理，确保安全性
2. **字段验证**: 必填字段不能为空，可选字段可以为空字符串
3. **状态管理**: 用户状态一旦设置，只能通过更新接口修改
4. **来源标识**: 用户来源在注册时设置，后续不可修改
5. **扩展信息**: extinfo 字段支持存储任意格式的扩展数据

---

## 更新日志

- **v1.0.0**: 初始版本，支持基础的用户管理功能
- 支持用户注册、登录、信息管理
- 支持JWT认证
- 支持用户状态和来源管理

---

## 技术支持

如有问题，请联系开发团队或查看服务日志获取详细信息。
