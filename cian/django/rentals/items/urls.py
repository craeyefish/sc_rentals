from django.urls import path

from . import views

urlpatterns = [
    path('', views.external_home, name='external_home'),
    path('home/', views.home, name='home'),
    path('login/', views.login, name='login'),
    path('signup/', views.signup, name='signup'),
]