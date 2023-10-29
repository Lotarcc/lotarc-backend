import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { MetricsInterceptor } from './metrics/metrics.interceptor';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('api');

  // Apply the metrics interceptor globally
  app.useGlobalInterceptors(new MetricsInterceptor());

  // Swagger
  const options = new DocumentBuilder()
    .setTitle('Lotarc Backend API')
    .setDescription('Lotarc Metrics API')
    .setVersion('1.0')
    .addTag('metrics')
    .build();

  const document = SwaggerModule.createDocument(app, options);
  SwaggerModule.setup('docs', app, document);
  await app.listen(3000);
}
bootstrap();
