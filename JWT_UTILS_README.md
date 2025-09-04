# JWT 验签工具 - TypeScript 版本

这是一个用于验证 Hulio User Service 生成的 JWT Token 的 TypeScript 工具库。

## 安装依赖

```bash
npm install jsonwebtoken
npm install --save-dev @types/jsonwebtoken
```

## 文件说明

- `jwt-utils.ts` - 核心验签工具文件

## 核心功能

### 1. 解析和验证 JWT Token

```typescript
import { parseToken, UserClaims } from './jwt-utils';

try {
  const claims: UserClaims = parseToken('your-jwt-token-here');
  console.log('用户ID:', claims.userId);
  console.log('用户名:', claims.username);
} catch (error) {
  console.error('Token验证失败:', error.message);
}
```

### 2. 检查 Token 是否有效

```typescript
import { isTokenValid } from './jwt-utils';

const isValid = isTokenValid('your-jwt-token-here');
if (isValid) {
  console.log('Token有效');
} else {
  console.log('Token无效');
}
```

### 3. 从 Authorization 头提取 Token

```typescript
import { extractTokenFromHeader } from './jwt-utils';

const authHeader = 'Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...';
const token = extractTokenFromHeader(authHeader);
if (token) {
  console.log('提取到的Token:', token);
}
```

### 4. 检查 Token 是否即将过期

```typescript
import { isTokenExpiringSoon } from './jwt-utils';

const token = 'your-jwt-token-here';
if (isTokenExpiringSoon(token)) {
  console.log('Token即将过期，建议刷新');
}
```

## 在 Next.js 中使用

### API 路由示例

```typescript
// app/api/protected/route.ts
import { NextRequest, NextResponse } from 'next/server';
import { parseToken, extractTokenFromHeader } from '../../../jwt-utils';

export async function GET(request: NextRequest) {
  try {
    // 从请求头获取 Authorization
    const authHeader = request.headers.get('authorization');
    
    if (!authHeader) {
      return NextResponse.json(
        { success: false, code: 401, message: '未授权访问' },
        { status: 401 }
      );
    }
    
    // 提取 Token
    const token = extractTokenFromHeader(authHeader);
    if (!token) {
      return NextResponse.json(
        { success: false, code: 401, message: 'Token格式错误' },
        { status: 401 }
      );
    }
    
    // 验证 Token
    const claims = parseToken(token);
    
    return NextResponse.json({
      success: true,
      code: 200,
      message: '验证成功',
      data: {
        userId: claims.userId,
        username: claims.username
      }
    });
    
  } catch (error) {
    return NextResponse.json(
      { success: false, code: 401, message: error.message },
      { status: 401 }
    );
  }
}
```

### 中间件示例

```typescript
// middleware.ts
import { NextRequest, NextResponse } from 'next/server';
import { isTokenValid, extractTokenFromHeader } from './jwt-utils';

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;
  
  // 只处理需要保护的路径
  if (!pathname.startsWith('/api/protected/')) {
    return NextResponse.next();
  }
  
  const authHeader = request.headers.get('authorization');
  const token = extractTokenFromHeader(authHeader || '');
  
  if (!token || !isTokenValid(token)) {
    return NextResponse.json(
      { success: false, code: 401, message: '未授权访问' },
      { status: 401 }
    );
  }
  
  return NextResponse.next();
}

export const config = {
  matcher: ['/api/protected/:path*']
};
```

## 在 Express.js 中使用

```typescript
import express from 'express';
import { parseToken, extractTokenFromHeader } from './jwt-utils';

const app = express();

// 认证中间件
function authMiddleware(req: express.Request, res: express.Response, next: express.NextFunction) {
  try {
    const authHeader = req.headers.authorization;
    const token = extractTokenFromHeader(authHeader || '');
    
    if (!token) {
      return res.status(401).json({
        success: false,
        code: 401,
        message: '未提供Token'
      });
    }
    
    const claims = parseToken(token);
    req.user = claims; // 将用户信息添加到请求对象
    next();
    
  } catch (error) {
    return res.status(401).json({
      success: false,
      code: 401,
      message: error.message
    });
  }
}

// 使用认证中间件
app.get('/api/protected', authMiddleware, (req, res) => {
  res.json({
    success: true,
    message: '访问成功',
    user: req.user
  });
});
```

## 环境配置

工具会根据 `NODE_ENV` 环境变量选择对应的公钥：

- `test` - 使用测试环境公钥
- `production` - 使用生产环境公钥
- 其他值 - 默认使用生产环境公钥

## 错误处理

工具会抛出以下类型的错误：

- `Token已过期` - Token 超过有效期
- `Token尚未生效` - Token 的生效时间未到
- `JWT验证失败` - Token 格式或签名验证失败
- `Token解析失败` - 其他解析错误

## 类型定义

```typescript
interface UserClaims {
  username: string;    // 用户名
  userId: number;      // 用户ID
  iat?: number;        // 签发时间
  exp?: number;        // 过期时间
  nbf?: number;        // 生效时间
  iss?: string;        // 签发者
  sub?: string;        // 主题
}
```

## 注意事项

1. **公钥安全**: 公钥已硬编码在工具中，确保与 Go 服务使用相同的密钥对
2. **环境区分**: 根据 `NODE_ENV` 自动选择对应的公钥
3. **错误处理**: 所有函数都有完善的错误处理机制
4. **类型安全**: 提供完整的 TypeScript 类型定义

## 与 Go 服务的兼容性

此工具完全兼容 Hulio User Service 的 JWT 实现：

- 使用相同的 RSA 密钥对
- 支持 RS256 算法
- 验证相同的 issuer 和 subject
- 解析相同的用户声明结构

## 使用建议

1. 在 API 路由中使用 `parseToken` 进行完整验证
2. 在中间件中使用 `isTokenValid` 进行快速检查
3. 使用 `isTokenExpiringSoon` 实现自动刷新机制
4. 始终进行错误处理，提供友好的错误信息
